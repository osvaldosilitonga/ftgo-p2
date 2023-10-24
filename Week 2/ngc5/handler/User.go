package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"ngc5/entitiy"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) User {
	return User{
		DB: db,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (handler User) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := entitiy.UserLogin{}

	// Parse body payload
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"msg":    "Json Decoder",
			"detail": err.Error(),
			"user":   r.Body,
		})
		return
	}

	isValid := LoginBodyValidation(w, body)
	if !isValid {
		return
	}

	query := `
		SELECT password, role FROM users
		WHERE email = ?;
	`

	var hash string
	row := handler.DB.QueryRowContext(ctx, query, body.Email)
	err = row.Scan(&hash, &body.Role)
	if err != nil {
		ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"msg":    "Email doesn't exist",
			"detail": err.Error(),
		})
		return
	}

	passMatch := checkPasswordHash(body.Password, hash)
	if !passMatch {
		ResponseJSON(w, http.StatusBadRequest, map[string]any{
			"msg": "Password not match",
		})
		return
	}

	res, token := GenerateToken(w, body)
	if !res {
		return
	}

	ResponseJSON(w, http.StatusBadRequest, map[string]any{
		"msg":          "Login Success",
		"access_token": token,
	})
}

func (handler User) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// user object
	user := entitiy.User{}

	// Parse body payload
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"msg":    "Json Decoder",
			"detail": err.Error(),
			"user":   r.Body,
		})
		return
	}

	isValid := BodyValidation(w, user)
	if !isValid {
		return
	}

	hashPassword, err := hashPassword(user.Password)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"msg":    "internal server error",
			"detail": err.Error(),
		})
		return
	}

	query := `
		INSERT INTO users (email, password, full_name, age, occupation)
		VALUES (?, ?, ?, ?, ?);	
	`

	// context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := handler.DB.ExecContext(ctx, query, user.Email, hashPassword, user.FullName, user.Age, user.Occupation)
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"detail":  err.Error(),
		})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"detail":  err.Error(),
		})
		return
	}

	user.ID = int(id)

	aff, err := result.RowsAffected()
	if err != nil {
		ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"detail":  err.Error(),
		})
		return
	}

	if aff == 0 {
		ResponseJSON(w, http.StatusInternalServerError, map[string]any{
			"message": "internal server error",
			"detail":  err.Error(),
		})
		return
	}

	ResponseJSON(w, http.StatusCreated, map[string]any{
		"message": "success create",
		"user":    user,
	})
}
