package order_service

import (
	"kominfo-assignment-2/dto"
	"kominfo-assignment-2/entity"
	"kominfo-assignment-2/repository/order_repo"
)

type Service interface {
	CreateOrderWithItems(payload dto.CreateOrderRequestDto) error
}

type orderService struct {
	orderRepo order_repo.Repository
}

func NewService(orderRepo order_repo.Repository) Service {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) CreateOrderWithItems(payload dto.CreateOrderRequestDto) error {

	order := entity.Order{
		OrderedAt:    payload.OrderedAt,
		CustomerName: payload.CustomerName,
	}

	items := []entity.Item{}

	for _, value := range payload.Items {
		item := entity.Item{
			ItemCode:    value.ItemCode,
			Quantity:    uint(value.Quantity),
			Description: value.Description,
		}

		items = append(items, item)
	}

	err := o.orderRepo.CreateOrderWithItems(order, items)

	if err != nil {
		return err
	}

	return nil
}
