package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Scrowszinho/api-go/internal/entity"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Test Product", 100.50)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
}

func TestFindProductById(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", rand.Float64()*100)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product, err = productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 1", product.Name)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", rand.Float64()*100)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product.Name = "Product 2"
	err = productDB.Update(product)
	assert.NoError(t, err)
	product, err = productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}
func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", rand.Float64()*100)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindById(product.ID.String())
	assert.Error(t, err)
}
