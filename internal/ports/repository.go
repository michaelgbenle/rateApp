package ports

import "github.com/michaelgbenle/rateApp/internal/models"

type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	TokenInBlacklist(token *string) bool
	AddTransaction(transaction *models.Transaction) error
	GetTransactions(user *models.User) (*[]models.Transaction, error)
	AddTokenToBlacklist(email string, token string) error
	UpdateUserbalances(user *models.User, exchange *models.Exchange, value float64) (*models.Transaction, error)
}
