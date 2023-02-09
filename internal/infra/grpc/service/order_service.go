package service

import (
	"context"

	"github.com/MrHenri/CleanArch/internal/infra/grpc/pb"
	"github.com/MrHenri/CleanArch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetOrderUseCase    usecase.GetOrderUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	getOrderUseCase usecase.GetOrderUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		GetOrderUseCase:    getOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.GetOrderRequest) (*pb.ListOrderResponse, error) {
	output, err := s.GetOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orderResponse []*pb.OrderResponse

	for _, order := range output {
		orderResponse = append(orderResponse, &pb.OrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.ListOrderResponse{Response: orderResponse}, nil
}
