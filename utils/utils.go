package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	model "prtvi/joke-api/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// create bson for joke to be inserted
func CreateJokeBson(joke model.Joke) bson.D {
	return bson.D{
		{Key: "id", Value: joke.ID},
		{Key: "type", Value: joke.Type},
		{Key: "setup", Value: joke.Setup},
		{Key: "punchline", Value: joke.Punchline},
	}
}

// helper function to load jokes to db using data.json file
func LoadJokes(collection mongo.Collection) {
	var jokes []model.Joke

	byteValue, _ := ioutil.ReadFile("data.json")
	json.Unmarshal(byteValue, &jokes)

	for i, joke := range jokes {
		joke.ID = i + 1
		jokeBson := CreateJokeBson(joke)

		_, err := collection.InsertOne(context.TODO(), jokeBson)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Println("Jokes loaded")
}

// get "n" random jokes from db
func GetJokes(n int, collection mongo.Collection) []model.Joke {
	pipeline := []bson.D{{{Key: "$sample", Value: bson.D{{Key: "size", Value: n}}}}}

	cursor, err := collection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		fmt.Println(err.Error())
	}

	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println(err.Error())
	}

	jokes := make([]model.Joke, len(results))

	for i, joke := range results {
		jokes[i] = MongoDocToStruct(joke)
	}

	return jokes
}

// convert mongo document to model.Joke struct (removes _id from mongo document)
func MongoDocToStruct(joke bson.M) model.Joke {
	jokeByte, _ := json.Marshal(joke)
	var jokeStruct model.Joke
	json.Unmarshal(jokeByte, &jokeStruct)

	return jokeStruct
}

// constructor for error responses
func NewError(code int, msg string) model.ErrResponse {
	return model.ErrResponse{StatusCode: code, Message: msg}
}

// helper function to check if request body is empty
func IfBodyContent(joke model.Joke, err error) bool {
	return err != nil || len(joke.Punchline) == 0 || len(joke.Setup) == 0 || len(joke.Type) == 0
}
