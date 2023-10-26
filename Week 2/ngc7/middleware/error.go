package middleware

import (
	"ngc7/utils"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				utils.ErrorMessage(c, &utils.ErrInternalServer)
			}
		}()

		c.Next()
	}
}
