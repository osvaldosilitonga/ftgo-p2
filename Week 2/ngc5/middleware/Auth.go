package middleware

import (
	"fmt"
	"net/http"
	"ngc5/handler"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

var secretKey = []byte("12345")

func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			handler.ResponseJSON(w, http.StatusUnauthorized, map[string]any{
				"msg": "Token Not Found",
			})
			return
		}

		// Method Parse
		parsedToken, parseErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Algoritma tidak valid")
			}

			return secretKey, nil
		})

		if parsedToken.Valid {
			fmt.Println(parsedToken)
		}

		if parseErr != nil || !parsedToken.Valid {
			fmt.Println("Error while decode token : ", parseErr)
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		next(w, r, p)
	}
}
