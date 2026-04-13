package repository

import (
	"payment-service/internal/domain"

	"github.com/jmoiron/sqlx"
)

type PaymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) Create(p *domain.Payment) error {
	query := `
        INSERT INTO payments (id, order_id, transaction_id, amount, status)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := r.db.Exec(query, p.ID, p.OrderID, p.TransactionID, p.Amount, p.Status)
	return err
}

func (r *PaymentRepository) GetByOrderID(orderID string) (*domain.Payment, error) {
	var p domain.Payment
	query := `SELECT * FROM payments WHERE order_id=$1`

	err := r.db.Get(&p, query, orderID)
	if err != nil {
		return nil, err
	}

	return &p, nil
}
