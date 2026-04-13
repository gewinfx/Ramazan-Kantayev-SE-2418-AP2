package usecase

import (
	"errors"
	"time"

	"order-service/internal/domain"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(order *domain.Order) error
	UpdateStatus(id string, status string) error
	GetByID(id string) (*domain.Order, error)
}

type PaymentClient interface {
	Pay(orderID string, amount int64) (string, error)
}

type OrderUsecase struct {
	repo          OrderRepository
	paymentClient PaymentClient
}

func NewOrderUsecase(repo OrderRepository, paymentClient PaymentClient) *OrderUsecase {
	return &OrderUsecase{
		repo:          repo,
		paymentClient: paymentClient,
	}
}

func (u *OrderUsecase) CreateOrder(customerID, itemName string, amount int64) (*domain.Order, error) {

	if amount <= 0 {
		return nil, errors.New("amount must be greater than 0")
	}

	order := &domain.Order{
		ID:         uuid.New().String(),
		CustomerID: customerID,
		ItemName:   itemName,
		Amount:     amount,
		Status:     "Pending",
		CreatedAt:  time.Now(),
	}

	err := u.repo.Create(order)
	if err != nil {
		return nil, err
	}

	status, err := u.paymentClient.Pay(order.ID, order.Amount)

	if err != nil {
		_ = u.repo.UpdateStatus(order.ID, "Failed")
		return nil, errors.New("payment service unavailable")
	}

	if status == "Authorized" {
		_ = u.repo.UpdateStatus(order.ID, "Paid")
		order.Status = "Paid"
	} else {
		_ = u.repo.UpdateStatus(order.ID, "Failed")
		order.Status = "Failed"
	}

	return order, nil
}

func (u *OrderUsecase) GetOrder(id string) (*domain.Order, error) {
	return u.repo.GetByID(id)
}

func (u *OrderUsecase) CancelOrder(id string) error {
	order, err := u.repo.GetByID(id)
	if err != nil {
		return err
	}

	if order.Status != "Pending" {
		return errors.New("only pending orders can be cancelled")
	}

	return u.repo.UpdateStatus(id, "Cancelled")
}
