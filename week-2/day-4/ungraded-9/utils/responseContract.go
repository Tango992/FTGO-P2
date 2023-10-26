package utils

import "net/http"

type ErrResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description any    `json:"description"`
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

	ErrDataNotFound = ErrResponse{
		Code:    http.StatusNotFound,
		Message: "Data not found",
	}

	ErrInternalServer = ErrResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
	}
)
