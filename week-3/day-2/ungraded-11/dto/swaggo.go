package dto

import "ungraded-11/entity"

type Error struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Message string `json:"message" extensions:"x-order=0"`
	Token   string `json:"token" extensions:"x-order=1"`
}

type RegisterResponse struct {
	ID            uint    `json:"id" extensions:"x-order=0"`
	Username      string  `json:"username" extensions:"x-order=1"`
	DepositAmount float32 `json:"deposit_amount" extensions:"x-order=2"`
}

type TransactionResponse struct {
	Message string `json:"message" extensions:"x-order=0"`
	entity.Transaction `json:"data" extensions:"x-order=1"`
}