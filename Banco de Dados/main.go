package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.NewString(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProduct("Nome", 10.50)
	err = insertProducts(db, product)

	product.Name = "Produto 1"

	updateProducts(db, product)

	if err != nil {
		panic(err)
	}

	productSelected, err := selectProduct(db, product.ID)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Product %v, possui o preco de %.2f \n", productSelected.Name, product.Price)

	products, err := selectAllProducts(db)

	if err != nil {
		panic(err)
	}

	for _, p := range *products {
		fmt.Printf("Product %v, possui o preco de %.2f \n", p.Name, p.Price)
	}

	err = deleteProduct(db, product.ID)

	if err != nil {
		panic(err)
	}

	products, err = selectAllProducts(db)

	if err != nil {
		panic(err)
	}

	for _, p := range *products {
		fmt.Printf("Product %v, foi mantido \n", p.Name)
	}
}

func insertProducts(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("INSERT INTO products (ID, NAME, PRICE) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.ID, product.Name, product.Price)

	return err
}

func updateProducts(db *sql.DB, product *Product) error {

	stmt, err := db.Prepare("update products set name = ?, price = ? where id =? ")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.Name, product.Price, "86a26c87-1226-4601-a6e5-721313855cc3")

	return err
}

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	var p Product

	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)

	return &p, err
}

func selectAllProducts(db *sql.DB) (*[]Product, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return &products, err
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return err
}
