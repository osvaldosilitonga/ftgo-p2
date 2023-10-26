package main

import (
	"log"
	"ngc7/config"
	"ngc7/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := config.DBConn() // db config

	// Gorm AutoMigrate
	// db.AutoMigrate(&entity.Store{}, &entity.Product{})

	userHandler := handler.NewUserHandler(db)
	productHandler := handler.NewProductHandler(db)

	r := gin.Default()

	r.POST("/users/login", userHandler.Login)
	r.POST("/users/register", userHandler.Register)

	products := r.Group("/products")
	{
		products.POST("")
		products.GET("", productHandler.GetAllProduct)
		products.GET("/:id", productHandler.GetProductById)
		products.PUT("/:id")
		products.DELETE("/:id")
	}

	r.Run(":8080")
}
