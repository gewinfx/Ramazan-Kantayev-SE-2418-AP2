package usecase

import (
	"payment-service/internal/domain"

	"github.com/google/uuid"
)

type PaymentRepository interface {
	Create(payment *domain.Payment) error
	GetByOrderID(orderID string) (*domain.Payment, error)
}

type PaymentUsecase struct {
	repo PaymentRepository
}

func NewPaymentUsecase(repo PaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{repo: repo}
}

func (u *PaymentUsecase) ProcessPayment(orderID string, amount int64) (*domain.Payment, error) {
	payment := &domain.Payment{
		ID:      uuid.New().String(),
		OrderID: orderID,
		Amount:  amount,
	}

	if amount > 100000 {
		payment.Status = "Declined"
	} else {
		payment.Status = "Authorized"
		payment.TransactionID = uuid.New().String()
	}

	err := u.repo.Create(payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (u *PaymentUsecase) GetPayment(orderID string) (*domain.Payment, error) {
	return u.repo.GetByOrderID(orderID)
}
