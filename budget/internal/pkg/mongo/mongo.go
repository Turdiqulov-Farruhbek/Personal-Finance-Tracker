package mong

import (
	"context"

	"finance_tracker/budget/internal/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New(cfg *config.Config) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + cfg.MongoUrl). //   DO NOT FORGET TO GET OUT OF COMMENT BEFORE DOCKER !!!!
											SetAuth(options.Credential{Username: "root", Password: "root"})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database("finance_tracker")

	return db, nil
}
