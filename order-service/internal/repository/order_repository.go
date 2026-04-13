package repository

import (
	"order-service/internal/domain"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// 💾 Создание заказа
func (r *OrderRepository) Create(order *domain.Order) error {
	query := `
		INSERT INTO orders (id, customer_id, item_name, amount, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		query,
		order.ID,
		order.CustomerID,
		order.ItemName,
		order.Amount,
		order.Status,
		order.CreatedAt,
	)

	return err
}

func (r *OrderRepository) UpdateStatus(id string, status string) error {
	query := `
		UPDATE orders
		SET status = $1
		WHERE id = $2
	`

	_, err := r.db.Exec(query, status, id)
	return err
}

func (r *OrderRepository) GetByID(id string) (*domain.Order, error) {
	var order domain.Order

	query := `
		SELECT id, customer_id, item_name, amount, status, created_at
		FROM orders
		WHERE id = $1
	`

	err := r.db.Get(&order, query, id)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
