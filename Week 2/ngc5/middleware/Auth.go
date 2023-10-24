package middleware

import (
	"fmt"
	"net/http"
	"ngc5/handler"
	"strings"

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

		// if parsedToken.Valid {
		// 	fmt.Println(parsedToken)
		// }

		path := strings.Split(r.URL.Path, "/")
		// fmt.Println(path[1], "<-------- Path")

		// claims := parsedToken.Claims.(jwt.MapClaims)
		// c := jwt.MapClaims(claims)
		claims := jwt.MapClaims(parsedToken.Claims.(jwt.MapClaims))

		if path[1] != claims["role"] {
			handler.ResponseJSON(w, http.StatusUnauthorized, map[string]any{
				"msg": "Access Not Allowed",
			})
			return
		}

		// fmt.Println(claims, "<-------- Map Claims")
		// fmt.Println(claims["role"], "<-------- Role")

		if parseErr != nil || !parsedToken.Valid {
			fmt.Println("Error while decode token : ", parseErr)
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}

		next(w, r, p)
	}
}
