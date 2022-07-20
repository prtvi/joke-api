package config

import (
	"context"
	"fmt"
	"os"

	// utils "prtvi/joke-api/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Global mongo collection connection
var Collection mongo.Collection

// establish connection & initialize global collection connection
func EstablishConnection() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DB_URL")))
	if err != nil {
		fmt.Println(err)
	}

	Collection = *client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	// enable to insert jokes into db using data.json file
	// utils.LoadJokes(Collection)
}
