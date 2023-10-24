package handler

import (
	"net/http"
	"ngc5/entitiy"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("12345")

func GenerateToken(w http.ResponseWriter, body entitiy.UserLogin) (bool, string) {
	claims := jwt.MapClaims{
		"email": body.Email,
		"role":  body.Role,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"msg": "Failed create token",
		})
		return false, tokenString
	}

	return true, tokenString
}
