package usecase

import (
	"github.com/MrHenri/CleanArch/internal/entity"
)

type GetOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrderUseCase(orderRepository entity.OrderRepositoryInterface) *GetOrderUseCase {
	return &GetOrderUseCase{OrderRepository: orderRepository}
}

func (g *GetOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := g.OrderRepository.Get()
	if err != nil {
		return nil, err
	}

	var outputDTOs []OrderOutputDTO

	for _, order := range orders {
		outputDTOs = append(outputDTOs, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return outputDTOs, nil

}
