package repository

import (
	"context"
	"github.com/michaelgbenle/rateApp/internal/ports"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Mongo struct {
	DB *mongo.Client
}

//NewDB create/returns a new instance of our Database
func NewDB(DB *mongo.Client) ports.Repository {
	return &Mongo{
		DB: DB,
	}
}

//Initialize opens the database, create tables if not created and populate it if its empty and returns a DB
func Initialize(dbURI string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURI))

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connection successful")
	return client, nil
}
