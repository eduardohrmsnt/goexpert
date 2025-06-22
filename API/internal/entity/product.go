package entity

import (
	"errors"
	"time"

	"github.com/eduardohrmsnt/go-expert-api/pkg/entity"
	"github.com/google/uuid"
)

var (
	ErrIDisRequired    = errors.New("ID is required")
	ErrInvalidId       = errors.New("Invalid id")
	ErrNameIsRequired  = errors.New("Name is required")
	ErrPriceIsRequired = errors.New("Price is required")
	ErrInvalidPrice    = errors.New("Invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()

	return product, err
}

func (p *Product) Validate() error {
	if p.ID == uuid.Nil {
		return ErrIDisRequired
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}

	if _, err := uuid.Parse(p.ID.String()); err != nil {
		return ErrInvalidId
	}

	if p.Price == 0 {
		return ErrPriceIsRequired
	}

	if p.Price < 0 {
		return ErrInvalidPrice
	}

	return nil
}
