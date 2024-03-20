package order_pg

import (
	"database/sql"
	"kominfo-assignment-2/entity"
	"kominfo-assignment-2/repository/order_repo"
)

type orderPG struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) order_repo.Repository {
	return &orderPG{
		db: db,
	}
}

func (o *orderPG) CreateOrderWithItems(order entity.Order, items []entity.Item) error {

	tx, err := o.db.Begin()

	if err != nil {
		return err
	}

	var orderId uint

	err = tx.QueryRow(
		create_new_order,
		order.OrderedAt,
		order.CustomerName,
	).Scan(&orderId)

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range items {
		_, err = tx.Exec(
			create_new_item,
			item.ItemCode,
			item.Description,
			item.Quantity,
			orderId,
		)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
