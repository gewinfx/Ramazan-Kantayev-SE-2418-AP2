package grpc

import (
	"context"

	"payment-service/internal/usecase"
	"project/proto/paymentpb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PaymentServer struct {
	paymentpb.UnimplementedPaymentServiceServer
	usecase *usecase.PaymentUsecase
}

func NewPaymentServer(u *usecase.PaymentUsecase) *PaymentServer {
	return &PaymentServer{usecase: u}
}

func (s *PaymentServer) ProcessPayment(ctx context.Context, req *paymentpb.PaymentRequest) (*paymentpb.PaymentResponse, error) {

	payment, err := s.usecase.ProcessPayment(req.OrderId, req.Amount)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &paymentpb.PaymentResponse{
		Status:        payment.Status,
		TransactionId: payment.TransactionID,
	}, nil
}
