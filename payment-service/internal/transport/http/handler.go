package http

import (
	"payment-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase *usecase.PaymentUsecase
}

func NewHandler(u *usecase.PaymentUsecase) *Handler {
	return &Handler{usecase: u}
}

type PaymentRequest struct {
	OrderID string `json:"order_id"`
	Amount  int64  `json:"amount"`
}

func (h *Handler) CreatePayment(c *gin.Context) {
	var req PaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	payment, err := h.usecase.ProcessPayment(req.OrderID, req.Amount)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, payment)
}

func (h *Handler) GetPayment(c *gin.Context) {
	orderID := c.Param("order_id")

	payment, err := h.usecase.GetPayment(orderID)
	if err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, payment)
}
