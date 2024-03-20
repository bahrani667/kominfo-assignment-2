package order_pg

const create_new_order = `
	INSERT INTO "orders"
	(ordered_at, customer_name)
	VALUES ($1, $2)
	RETURNING order_id
`

const create_new_item = `
	INSERT INTO "items"
	(item_code, description, quantity, order_id)
	VALUES($1, $2, $3, $4)
`
