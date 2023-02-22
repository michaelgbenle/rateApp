package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Email     string             `json:"email" binding:"required"`
	Password  string             `json:"password" binding:"required"`
	Balance   map[string]float64 `json:"balance"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

type Transaction struct {
	UserEmail       string             `json:"user_email"`
	Balance         map[string]float64 `json:"balance"`
	TransactionType string             `json:"transaction_type"`
	Success         bool               `json:"success"`
	CreatedAt       time.Time          `json:"created_at"`
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
