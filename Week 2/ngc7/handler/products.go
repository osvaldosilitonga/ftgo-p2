package handler

import (
	// "database/sql"

	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProductHandler(db *gorm.DB) Product {
	return Product{
		DB: db,
	}
}

func (handler Product) GetAllProduct(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"status":  http.StatusOK,
	})
}
func (handler Product) GetProductById(ctx *gin.Context) {}
