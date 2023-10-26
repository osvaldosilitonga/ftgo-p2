package handler

import (
	// "database/sql"

	"net/http"
	"ngc7/entity"
	"ngc7/utils"

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
	products := []entity.Product{}

	result := handler.DB.Find(&products)
	if result.Error != nil {
		utils.ErrorMessage(ctx, &utils.ErrInternalServer)
		return
	}

	if result.RowsAffected == 0 {
		utils.ErrorMessage(ctx, &utils.ErrDataNotFound)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "OK",
		"products": products,
	})
}

func (handler Product) GetProductById(ctx *gin.Context) {}

func (handler Product) PostProduct(ctx *gin.Context) {
	product := entity.Product{}

	// request body
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		utils.ErrorMessage(ctx, &utils.ErrBadRequest)
		return
	}

	r := handler.DB.Preload("Product").Create(&product)
	if r.Error != nil {
		utils.ErrorMessage(ctx, &utils.ErrInternalServer)
		return
	}

	if r.RowsAffected == 0 {
		utils.ErrorMessage(ctx, &utils.ErrBadRequest)
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Created Success",
		"data":    product,
	})
}
