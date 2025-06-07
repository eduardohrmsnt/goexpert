package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := &Category{Name: "Eletronicos"}

	db.Create(category)

	db.Create(&Product{
		Name:  "Notebook",
		Price: 10.60,
		Categories: []Category{
			*category,
		},
	})

	var products []Product
	db.Preload("Categories").Find(&products)

	/* 	products := &[]Product{
	   		{Name: "Sou produto", Price: 10.0},
	   	}

	   	db.Create(products)

	   	createProduct(db, &Product{
	   		Name:  "Produtao",
	   		Price: 10.5,
	   	}) */
	/*
		var productSearch Product

		db.First(&productSearch, 2)

		db.First(&productSearch, "name = ?", "Produtao")

		fmt.Println(productSearch)
	*/
	/* 	var productsSearch []Product

	   	db.Limit(2).Offset(2).Find(&productsSearch)

	   	for _, products := range productsSearch {
	   		fmt.Println(products)
	   	} */
	/* 	var p Product

	   	db.First(&p)

	   	p.Name = "Mouse"

	   	db.Save(&p)

	   	var p2 Product

	   	db.Where("id = ?", &p.ID).First(&p2)

	   	fmt.Println(p2)

	   	db.Delete(&p2)

	   	var p3 Product

	   	db.Where("id = ?", &p.ID).FirstOrInit(&p3)

	   	fmt.Println(p3) */
}

func createProduct(db *gorm.DB, product *Product) {
	db.Create(product)
}
