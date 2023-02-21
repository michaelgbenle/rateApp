package models

import "time"

type User struct {
	ID        int                `json:"id"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Balance   map[string]float64 `json:"balance"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type Transaction struct {
	ID        int                `json:"id"`
	UserEmail string             `json:"user_email"`
	Balance   map[string]float64 `json:"balance"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type Blacklist struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
}
