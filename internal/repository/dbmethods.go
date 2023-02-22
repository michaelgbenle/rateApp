package repository

import (
	"context"
	"errors"
	"github.com/michaelgbenle/rateApp/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

func (m *Mongo) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := m.DB.Database("payourse").Collection("users").FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, err
}
func (m *Mongo) CreateUser(user *models.User) error {
	res, err := m.DB.Database("payourse").Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	log.Println(res.InsertedID)
	return nil
}
func (m *Mongo) TokenInBlacklist(token *string) bool {
	res := m.DB.Database("token").Collection("blacklist").FindOne(context.Background(), bson.M{"token": token})
	if res.Err() != nil {
		return true
	}
	return false
}
func (m *Mongo) AddTransaction(transaction *models.Transaction) error {
	res, err := m.DB.Database("payourse").Collection("transactions").InsertOne(context.Background(), transaction)
	if err != nil {
		return err
	}
	log.Println(res.InsertedID)
	return nil
}

func (m *Mongo) GetTransactions(user *models.User) (*[]models.Transaction, error) {

	filter := bson.D{{"email", user.Email}}

	cur, err := m.DB.Database("payourse").Collection("transactions").Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	var transactions []models.Transaction
	if err = cur.All(context.Background(), &transactions); err != nil {
		return nil, err
	}
	return &transactions, err
}
func (m *Mongo) AddTokenToBlacklist(email string, token string) error {
	res, err := m.DB.Database("token").Collection("blacklist").InsertOne(context.Background(), bson.M{"email": email, "token": token})
	if err != nil {
		return err
	}
	log.Println(res.InsertedID)
	return nil
}

func (m *Mongo) UpdateUserbalances(user *models.User, exchange *models.Exchange, value float64) (*models.Transaction, error) {

	if exchange.Currency == ("NGN") {
		filter := bson.D{{"email", user.Email}}
		update1 := bson.D{
			{"$set", bson.D{
				{"balance", user.Balance[string(exchange.Currency)] - exchange.Amount}, {"updated_at", time.Now()}}},
		}
		update2 := bson.D{
			{"$set", bson.D{
				{"balance", user.Balance["USD"] + value}, {"updated_at", time.Now()}}},
		}
		_, err := m.DB.Database("payourse").Collection("users").UpdateOne(context.Background(), filter, update1)
		if err != nil {
			return &models.Transaction{}, err
		}
		_, err = m.DB.Database("payourse").Collection("users").UpdateOne(context.Background(), filter, update2)
		if err != nil {
			return &models.Transaction{}, err
		}
		transaction := &models.Transaction{
			UserEmail:       user.Email,
			Balance:         user.Balance,
			TransactionType: "NGN to USD",
			Success:         true,
			CreatedAt:       time.Now(),
		}
		err = m.AddTransaction(transaction)
		if err != nil {
			return &models.Transaction{}, err
		}
		return transaction, nil
	}

	if exchange.Currency == ("USD") {
		filter := bson.D{{"email", user.Email}}
		update1 := bson.D{
			{"$set", bson.D{
				{"balance", user.Balance[string(exchange.Currency)] - exchange.Amount}, {"updated_at", time.Now()}}},
		}
		update2 := bson.D{
			{"$set", bson.D{
				{"balance", user.Balance["NGN"] + value}, {"updated_at", time.Now()}}},
		}
		_, err := m.DB.Database("payourse").Collection("users").UpdateOne(context.Background(), filter, update1)
		if err != nil {
			return &models.Transaction{}, err
		}
		_, err = m.DB.Database("payourse").Collection("users").UpdateOne(context.Background(), filter, update2)
		if err != nil {
			return &models.Transaction{}, err
		}
		transaction := &models.Transaction{
			UserEmail:       user.Email,
			Balance:         user.Balance,
			TransactionType: "USD to NGN",
			Success:         true,
			CreatedAt:       time.Now(),
		}
		err = m.AddTransaction(transaction)
		if err != nil {
			return &models.Transaction{}, err
		}
		return transaction, nil
	}

	return &models.Transaction{}, nil
}
