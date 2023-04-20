package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 11.5)
	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("1", 0)
	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("1", -10)
	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Bola", 10.50)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Bola", product.Name)
	assert.Equal(t, 10.50, product.Price)
}

func TestValidateProduct(t *testing.T) {
	product, err := NewProduct("Bola", 10.50)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Nil(t, product.Validate())
}
