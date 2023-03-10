package models

import (
	"time"
)

type User struct {
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Balance   Balance   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Balance struct {
	NGN float64 `json:"NGN"`
	USD float64 `json:"USD"`
}

type Transaction struct {
	UserEmail       string    `json:"user_email"`
	Balance         Balance   `json:"balance"`
	TransactionType string    `json:"transaction_type"`
	Success         bool      `json:"success"`
	CreatedAt       time.Time `json:"created_at"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Blacklist struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}
