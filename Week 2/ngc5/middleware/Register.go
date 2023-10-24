package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ngc5/entitiy"
	"ngc5/handler"
	"reflect"
	"regexp"

	"github.com/julienschmidt/httprouter"
)

func isRequired(user entitiy.User) (bool, string) {
	t := reflect.TypeOf(user)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(user).Field(i).Interface()
			if value == "" {
				return false, fmt.Sprintf("%v is required", t.Field(i).Tag.Get("json"))
			}
		}
	}
	return true, ""
}

func emaiValidation(email string) bool {
	emaiRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emaiRegex.MatchString(email)
}

func BodyCheck(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		// user object
		user := entitiy.User{}

		// Parse body payload
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			handler.ResponseJSON(w, http.StatusInternalServerError, map[string]any{
				"msg":    "Internal Server Error",
				"detail": err.Error(),
			})
			return
		}

		// Required Validation
		res, msg := isRequired(user)
		if res == false {
			handler.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": msg,
			})
			return
		}

		// Email validation
		isValid := emaiValidation(user.Email)
		if !isValid {
			handler.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "Email not valid",
			})
			return
		}

		// Password Validation
		if len(user.Password) < 8 {
			handler.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "Password to short, min 8 character",
			})
			return
		}

		// Full Name Validation
		if len(user.FullName) < 6 || len(user.FullName) > 15 {
			handler.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "Full name must be 6 - 15 character",
			})
			return
		}

		// Age Validation
		if user.Age < 17 {
			handler.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "Age minimal 17 years old",
			})
			return
		}

		next(w, r, p)
	}
}
