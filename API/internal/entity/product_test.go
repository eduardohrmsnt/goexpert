package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("produto 1", 10)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "produto 1", product.Name)
	assert.Equal(t, 10.0, product.Price)
}

func TestProductWithInvalidName(t *testing.T) {
	_, err := NewProduct("", 10)

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrNameIsRequired)
}

func TestProductWithRequiredPrice(t *testing.T) {
	_, err := NewProduct("produto 1", 0)

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrPriceIsRequired)
}

func TestProductWithInvalidPrice(t *testing.T) {
	_, err := NewProduct("produto 1", -1)

	assert.NotNil(t, err)
	assert.Equal(t, err, ErrInvalidPrice)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("produto 1", 10)

	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
