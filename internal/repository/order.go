package repository

import (
	"time"

	"github.com/Anwarjondev/fast-food/internal/db"
)

type Order struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	Status      string    `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	DeliveredAt time.Time `json:"delivered_at" db:"delivered_at"`
	TotalPrice  float64   `json:"total_price" db:"total_amount"`
}
type OrderDetail struct {
	FoodID int `json:"food_id"`
	Count  int `json:"count"`
}

func CreateOrder(UserID int, fooditems []OrderDetail) (int, error) {
	var orderID int
	var totalOrderPrice float64
	tx := db.DB.MustBegin()
	err := tx.QueryRow(`Insert into orders(user_id, total_amount, created_at, status) values($1, 0, now(), 'active') returning id`, UserID).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	for _, item := range fooditems {
		var price float64
		err = tx.Get(&price, `SELECT price FROM food WHERE id = $1`, item.FoodID)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		itemTotal := price * float64(item.Count)
		totalOrderPrice += itemTotal
		_, err = tx.Exec(`Insert into order_detail (order_id, food_id, count) values($1, $2, $3)`, orderID, item.FoodID, item.Count)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	_, err = tx.Exec(`UPDATE orders SET total_amount = $1 WHERE id = $2`, totalOrderPrice, orderID)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	err = tx.Commit()
	return orderID, err
}
func GetAllOrderByStatus(UserID int, status string) ([]Order, error) {
	var orders []Order
	var err error
	if status == "all" {
		err = db.DB.Select(&orders, `select *from orders where user_id = $1`, UserID)
	} else {
		err = db.DB.Select(&orders, `select *from orders where user_id = $1 and status = $2`, UserID, status)
	}
	return orders, err
}
func CancelOrder(UserID, OrderID int) error {
	_, err := db.DB.Exec(`update orders set status = 'canceled' where id = $1 and user_id = $2 and status = 'active' and now() - created_at < interval '10 minutes'`, OrderID, UserID)
	return err
}
