package handler

import (
	"fmt"
	"net/http"
	"ngc5/entitiy"
	"reflect"
	"regexp"
)

func isRequired(user interface{}) (bool, string) {
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

func LoginBodyValidation(w http.ResponseWriter, body entitiy.UserLogin) bool {
	// Required Validation
	res, msg := isRequired(body)
	if res == false {
		ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"message": msg,
		})
		return false
	}

	return true
}

func BodyValidation(w http.ResponseWriter, user entitiy.User) bool {
	// Required Validation
	res, msg := isRequired(user)
	if res == false {
		ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"message": msg,
		})
		return false
	}

	// Email validation
	isValid := emaiValidation(user.Email)
	if !isValid {
		ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"message": "Email not valid",
		})
		return false
	}

	// Password Validation
	if len(user.Password) < 8 {
		ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"message": "Password to short, min 8 character",
		})
		return false
	}

	// Full Name Validation
	if len(user.FullName) < 6 || len(user.FullName) > 15 {
		ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"message": "Full name must be 6 - 15 character",
		})
		return false
	}

	// Age Validation
	if user.Age < 17 {
		ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"message": "Age minimal 17 years old",
		})
		return false
	}

	return true
}
