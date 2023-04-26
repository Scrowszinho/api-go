package main

import (
	"net/http"

	"github.com/Scrowszinho/api-go/configs"
	"github.com/Scrowszinho/api-go/internal/entity"
	"github.com/Scrowszinho/api-go/internal/infra/database"
	"github.com/Scrowszinho/api-go/internal/webserver/handlers"
	"github.com/glebarez/sqlite"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Get("/products", productHandler.GetProducts)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB, config.TokenAuth, config.JWTExpiresIn)
	r.Post("/users", userHandler.CreateUser)
	r.Post("/users/login", userHandler.GetJWT)

	http.ListenAndServe(":8000", r)
}
