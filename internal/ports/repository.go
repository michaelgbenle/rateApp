package ports

import "github.com/michaelgbenle/rateApp/internal/models"

type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	TokenInBlacklist(token *string) bool
	GetTransactions(user *models.User) (*[]models.Transaction, error)
	AddTokenToBlacklist(email string, token string) error
}
