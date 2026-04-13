package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"order-service/internal/repository"
	httpHandler "order-service/internal/transport/http"
	"order-service/internal/usecase"
)

func main() {
	db, err := sqlx.Connect(
		"postgres",
		"postgres://postgres:postgres@localhost:5432/order_db?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewOrderRepository(db)

	paymentClient := usecase.NewPaymentClient()

	orderUsecase := usecase.NewOrderUsecase(repo, paymentClient)

	handler := httpHandler.NewHandler(orderUsecase)

	r := gin.Default()

	r.POST("/orders", handler.CreateOrder)
	r.GET("/orders/:id", handler.GetOrder)
	r.PATCH("/orders/:id/cancel", handler.CancelOrder)

	r.Run(":8081")
}
