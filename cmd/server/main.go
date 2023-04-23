package main

import (
	"encoding/json"
	"net/http"

	"github.com/Scrowszinho/api-go/configs"
	"github.com/Scrowszinho/api-go/internal/dto"
	"github.com/Scrowszinho/api-go/internal/entity"
	"github.com/Scrowszinho/api-go/internal/infra/database"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		println(2)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		println(3)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
