package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type HTTPPaymentClient struct {
	client *http.Client
}

func NewPaymentClient() *HTTPPaymentClient {
	return &HTTPPaymentClient{
		client: &http.Client{
			Timeout: 2 * time.Second,
		},
	}
}

type paymentRequest struct {
	OrderID string `json:"order_id"`
	Amount  int64  `json:"amount"`
}

type paymentResponse struct {
	Status string `json:"status"`
}

func (c *HTTPPaymentClient) Pay(orderID string, amount int64) (string, error) {

	reqBody := paymentRequest{
		OrderID: orderID,
		Amount:  amount,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Post(
		"http://localhost:8082/payments",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("payment service returned non-200 status")
	}

	var result paymentResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Status, nil
}
