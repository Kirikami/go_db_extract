package database

type User struct {
	UserID int    `db:"user_id"`
	Name   string `db:"name"`
}

type Seller struct {
	OrderID     int     `db:"order_id"`
	UserID      int     `db:"user_id"`
	OrderAmount float64 `db:"order_amount"`
}
