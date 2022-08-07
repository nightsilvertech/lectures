package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type Order struct {
	Id        int64
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetOrderByID(db *sql.DB, orderId int64) (*Order, error) {
	stmt, err := db.Prepare(`SELECT * FROM orders WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows := stmt.QueryRow(orderId)
	var order Order
	err = rows.Scan(
		&order.Id,
		&order.Username,
		&order.CreatedAt,
		&order.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func GetOrders(db *sql.DB) ([]Order, error) {
	stmt, err := db.Prepare(`SELECT * FROM orders`)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var orders []Order
	for rows.Next() {
		var order Order
		err = rows.Scan(
			&order.Id,
			&order.Username,
			&order.CreatedAt,
			&order.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func main() {
	connectionString := fmt.Sprintf("root:root@tcp(localhost:3306)/ecommers?parseTime=true")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	order, err := GetOrderByID(db, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(order)

	//orders, err := GetOrders(db)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, order := range orders {
	//	fmt.Println(order)
	//}
}
