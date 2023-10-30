package utils

import "net/http"

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details"`
}

var (
	ErrBadRequest = ErrResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad request",
	}

	ErrUnauthorized = ErrResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized access",
	}

	ErrNotFound = ErrResponse{
		Code:    http.StatusNotFound,
		Message: "Request data is not found",
	}

	ErrInternalServer = ErrResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}
)
