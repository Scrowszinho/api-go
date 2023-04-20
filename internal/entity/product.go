package entity

import (
	"errors"
	"time"

	"github.com/Scrowszinho/api-go/pkg/entity"
)

var (
	ErrIdIsRequired    = errors.New("id is required")
	ErrInvalidId       = errors.New("id is invalid")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0.0 {
		return ErrPriceIsRequired
	}
	if p.ID.String() == "" {
		return ErrIdIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidId
	}
	if p.Price < 0.0 {
		return ErrInvalidPrice
	}
	return nil
}
