package dto

import "time"

type CreateOrderRequestDto struct {
	OrderedAt    time.Time              `json:"orderedAt"`
	CustomerName string                 `json:"customerName"`
	Items        []CreateItemRequestDto `json:"items"`
}
