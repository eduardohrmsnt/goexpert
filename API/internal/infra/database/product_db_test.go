package database

import (
	"fmt"
	"math/rand/v2"
	"testing"

	"github.com/eduardohrmsnt/go-expert-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
		return
	}

	db.AutoMigrate(&entity.Product{})

	productDb := NewProduct(db)

	product, err := entity.NewProduct("Meu produto", 10)

	if err != nil {
		t.Error(err)
		return
	}

	err = productDb.Create(product)

	if err != nil {
		t.Error(err)
		return
	}

	var productFound entity.Product

	productDb.DB.Where("id = ?", product.ID.String()).First(&productFound)

	assert.Nil(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, productFound.Name, "Meu produto")
	assert.Equal(t, productFound.Price, 10)
	assert.Nil(t, product.Validate())
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
		return
	}

	db.AutoMigrate(&entity.Product{})

	productDb := NewProduct(db)

	product, err := entity.NewProduct("Meu produto", 10)

	if err != nil {
		t.Error(err)
		return
	}

	err = productDb.Create(product)

	if err != nil {
		t.Error(err)
		return
	}

	productFound, err := productDb.FindByID(product.ID.String())

	if err != nil {
		t.Error(err)
		return
	}

	assert.Nil(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, productFound.Name, "Meu produto")
	assert.Equal(t, productFound.Price, 10.0)
	assert.Nil(t, product.Validate())
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
		return
	}

	db.AutoMigrate(&entity.Product{})

	productDb := NewProduct(db)

	product, err := entity.NewProduct("Meu produto", 10)

	if err != nil {
		t.Error(err)
		return
	}

	err = productDb.Create(product)

	if err != nil {
		t.Error(err)
		return
	}

	product.Price = 30

	err = productDb.Update(product)

	if err != nil {
		t.Error(err)
		return
	}

	var productFound entity.Product

	productDb.DB.Where("id = ?", product.ID.String()).First(&productFound)

	assert.Nil(t, err)
	assert.NotNil(t, productFound)
	assert.Equal(t, productFound.Name, "Meu produto")
	assert.Equal(t, productFound.Price, 30.0)
	assert.Nil(t, product.Validate())
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
		return
	}

	db.AutoMigrate(&entity.Product{})

	productDb := NewProduct(db)

	product, err := entity.NewProduct("Meu produto", 10)

	if err != nil {
		t.Error(err)
		return
	}

	err = productDb.Create(product)

	if err != nil {
		t.Error(err)
		return
	}

	err = productDb.Delete(product.ID.String())

	if err != nil {
		t.Error(err)
		return
	}

	_, err = productDb.FindByID(product.ID.String())

	assert.Error(t, err)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})

	if err != nil {
		t.Error(err)
		return
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Produto %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDb := NewProduct(db)

	products, err := productDb.FindAll(1, 10, "asc")

	assert.NoError(t, err)
	assert.Equal(t, "Produto 1", products[0].Name)
	assert.Equal(t, "Produto 10", products[9].Name)

	products, err = productDb.FindAll(2, 10, "asc")

	assert.NoError(t, err)
	assert.Equal(t, "Produto 11", products[0].Name)
	assert.Equal(t, "Produto 20", products[9].Name)

	products, err = productDb.FindAll(3, 10, "asc")

	assert.NoError(t, err)
	assert.Equal(t, 3, len(products))
	assert.Equal(t, "Produto 21", products[0].Name)
	assert.Equal(t, "Produto 23", products[2].Name)
}
