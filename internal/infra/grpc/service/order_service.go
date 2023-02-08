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

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) GetOrder(context.Context, *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	output, err := s.GetOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orderResponse []*pb.CreateOrderResponse

	for _, order := range output {
		orderResponse = append(orderResponse, &pb.CreateOrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.GetOrderResponse{Response: orderResponse}, nil
}