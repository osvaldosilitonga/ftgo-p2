package handler

import (
	// "database/sql"

	"net/http"
	"ngc7/dto"
	"ngc7/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) User {
	return User{
		DB: db,
	}
}

func (handler User) Login(ctx *gin.Context) {
	body := dto.StoreLogin{}
	store := entity.Store{}

	// request body
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"detail":  err.Error(),
		})
		return
	}

	// Check if returns RecordNotFound error
	if result := handler.DB.Where("store_email = ?", body.StoreEmail).First(&store); result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"detail":  result.Error.Error(),
		})
		return
	}

	if body.Password != store.Password {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Password not match",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Login Success",
	})
}

func (handler User) Register(ctx *gin.Context) {
	store := entity.Store{}

	// request body
	if err := ctx.ShouldBindJSON(&store); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
			"detail":  err.Error(),
		})
		return
	}

	r := handler.DB.Preload("Product").Create(&store)
	if r.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal Server Error",
			"detail":  r.Error,
		})
		return
	}

	if r.RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal Server Error",
			"detail":  r.Error,
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Created",
		"store":   store,
	})
}
