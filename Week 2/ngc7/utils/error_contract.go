package utils

import "net/http"

type APIError struct {
	Code    int
	Message string
}

var (
	ErrDataNotFound = APIError{
		Code:    http.StatusNotFound,
		Message: "Data Not Found",
	}
	ErrInternalServer = APIError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	ErrBadRequest = APIError{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
	}
	ErrInvalidCredential = APIError{
		Code:    http.StatusUnauthorized,
		Message: "Invalid Credential",
	}
	ErrNotAuthenticated = APIError{
		Code:    http.StatusUnauthorized,
		Message: "Not Authenticated",
	}
	ErrWrongPassword = APIError{
		Code:    http.StatusBadRequest,
		Message: "Wrong password",
	}
)
