package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// What is repository
// Layer that handles communication between database
// or third party (API, Library)
// or source of data (manipulate or retrieval)

type Order struct {
	Id        int64
	Username  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type repository struct {
	Database *sql.DB
}

type RepositoryOrder interface {
	// GetOrders query many orders
	GetOrders() ([]Order, error)
	// GetOrderByID query one order by its id
	GetOrderByID(orderId int64) (*Order, error)
	// WriteOrder insert new order to database
	WriteOrder(order Order) error
	// UpdateOrder update existing order from table order
	UpdateOrder(order Order) error
	// RemoveOrderByID remove order by id from table order
	RemoveOrderByID(id int64) error
}

func NewRepositoryOrder(db *sql.DB) RepositoryOrder {
	return &repository{
		Database: db,
	}
}

// WriteOrder insert new order to database
func (r *repository) WriteOrder(order Order) error {
	stmt, err := r.Database.Prepare("INSERT INTO orders (username) VALUE (?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(order.Username)
	if err != nil {
		return err
	}

	newRowInserted, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if newRowInserted != 0 {
		return nil
	} else {
		return errors.New("No changes inserted")
	}
}

// UpdateOrder update existing order from table order
func (r *repository) UpdateOrder(order Order) error {
	stmt, err := r.Database.Prepare("UPDATE orders SET username = ? WHERE id = ?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(order.Username, order.Id)
	if err != nil {
		return err
	}

	newRowInserted, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if newRowInserted != 0 {
		return nil
	} else {
		return errors.New("No changes updated")
	}
}

// RemoveOrderByID remove order by id from table order
func (r *repository) RemoveOrderByID(id int64) error {
	stmt, err := r.Database.Prepare("DELETE FROM orders WHERE id = ?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	newRowInserted, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if newRowInserted != 0 {
		return nil
	} else {
		return errors.New("No changes deleted")
	}
}

// GetOrders query many orders
func (r *repository) GetOrders() ([]Order, error) {
	stmt, err := r.Database.Prepare(`SELECT * FROM orders`)
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

// GetOrderByID query one order by its id
func (r *repository) GetOrderByID(orderId int64) (*Order, error) {
	stmt, err := r.Database.Prepare(`SELECT * FROM orders WHERE id = ?`)
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

	repo := NewRepositoryOrder(db)

	err = repo.RemoveOrderByID(6)
	if err != nil {
		log.Fatal(err)
	}

	orders, err := repo.GetOrders()
	if err != nil {
		log.Fatal(err)
	}

	for _, order := range orders {
		fmt.Println(order)
	}

	//orders, err := GetOrders(Database)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, order := range orders {
	//	fmt.Println(order)
	//}
}
