package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	repository "payment-service/internal/repository/postgres"
	"payment-service/internal/transport/http"
	"payment-service/internal/usecase"
)

func main() {
	db, err := sqlx.Connect("postgres", "postgres://postgres:postgres@localhost:5432/payment_db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewPaymentRepository(db)
	usecase := usecase.NewPaymentUsecase(repo)
	handler := http.NewHandler(usecase)

	r := gin.Default()

	r.POST("/payments", handler.CreatePayment)
	r.GET("/payments/:order_id", handler.GetPayment)

	r.Run(":8082")
}
