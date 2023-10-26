package handler

import (
	// "database/sql"

	"net/http"
	"ngc7/dto"
	"ngc7/entity"
	"ngc7/utils"

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

	// Find by store_email
	if result := handler.DB.Where("store_email = ?", body.StoreEmail).First(&store); result.Error != nil {
		utils.ErrorMessage(ctx, &utils.ErrDataNotFound)
		return
	}

	if body.Password != store.Password {
		utils.ErrorMessage(ctx, &utils.ErrWrongPassword)
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
		utils.ErrorMessage(ctx, &utils.ErrBadRequest)
		return
	}

	r := handler.DB.Preload("Product").Create(&store)
	if r.Error != nil {
		utils.ErrorMessage(ctx, &utils.ErrInternalServer)
		return
	}

	if r.RowsAffected == 0 {
		utils.ErrorMessage(ctx, &utils.ErrBadRequest)
		return
	}

	ctx.JSON(201, gin.H{
		"message": "Created",
		"store":   store,
	})
}
