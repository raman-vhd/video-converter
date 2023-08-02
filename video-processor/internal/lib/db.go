package lib

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const timeout = 10 * time.Second

type Database struct {
	MongoClient *mongo.Client
}

func NewDB(env Env) Database {
    uri := fmt.Sprintf("mongodb://%v:%v", env.DBHost, env.DBPort)
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("failed connecting to database: %v\n", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("failed pinging database: %v\n", err)
	}

	log.Println("Connected to MongoDB!")
	return Database{
		MongoClient: client,
	}
}

func (d Database) GetCollection(collectionName string) *mongo.Collection {
	collection := d.MongoClient.Database("vidAPI").Collection(collectionName)
	return collection
}
