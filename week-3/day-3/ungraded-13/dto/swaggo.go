package dto

import (
	"ungraded-13/entity"
)

type Error struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	Message string `json:"message" extensions:"x-order=0"`
	Token   string `json:"token" extensions:"x-order=1"`
}

type RegisterResponseTmp struct {
	ID            uint    `json:"id" extensions:"x-order=0"`
	Username      string  `json:"username" extensions:"x-order=1"`
	DepositAmount float32 `json:"deposit_amount" extensions:"x-order=2"`
}

type RegisterResponse struct {
	Message string              `json:"message" extensions:"x-order=0"`
	Data    RegisterResponseTmp `json:"data" extensions:"x-order=1"`
}

type TransactionResponse struct {
	Message string             `json:"message" extensions:"x-order=0"`
	Data    entity.Transaction `json:"data" extensions:"x-order=1"`
}

type StoreByIdResponse struct {
	Message           string `json:"message" extensions:"x-order=0"`
	StoreDetailSwaggo `json:"data" extensions:"x-order=1"`
}

type StoreDetailSwaggo struct {
	StoreWithSales `json:"store"`
	Weather        struct{} `json:"weather"`
}

type ProductResponse struct {
	Message string           `json:"message" extensions:"x-order=0"`
	Data    []entity.Product `json:"data" extensions:"x-order=1"`
}

type Store struct {
	ID      uint   `json:"store_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type StoreResponse struct {
	Message string `json:"message" extensions:"x-order=0"`
	Data    []Store  `json:"data" extensions:"x-order=1"`
}
