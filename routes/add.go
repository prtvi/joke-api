package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	config "jokeapi/config"
	model "jokeapi/model"
	utils "jokeapi/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// "/new"
// add new joke to collection

func AddOne(c echo.Context) error {
	fmt.Println("Endpoint hit: /new")

	// read request body
	var newJoke model.Joke
	err := json.NewDecoder(c.Request().Body).Decode(&newJoke)

	// if request body is empty
	if utils.IfBodyContent(newJoke, err) {
		err := utils.NewError(http.StatusBadRequest, "No request body found")
		return c.JSON(http.StatusBadRequest, err)
	}

	// check if joke exists in db (by punchline && setup)
	cursor, err := config.Collection.Find(context.TODO(), bson.D{
		{Key: "punchline", Value: newJoke.Punchline},
		{Key: "setup", Value: newJoke.Setup},
	})

	if err != nil {
		fmt.Println("err")
	}

	// read fetched docs from db
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		err := utils.NewError(http.StatusBadRequest, "Joke already exists in db")
		return c.JSON(http.StatusBadRequest, err)
	}

	// if joke exists in db then return bad request
	if len(results) >= 1 {
		err := utils.NewError(http.StatusBadRequest, "Joke already exists in db")
		return c.JSON(http.StatusBadRequest, err)
	}

	// else add new id and insert into db
	count, _ := config.Collection.CountDocuments(context.TODO(), bson.D{})

	// new id = len(documents) + 1
	newJoke.ID = int(count) + 1
	newJokeBson := utils.CreateJokeBson(newJoke)

	_, err = config.Collection.InsertOne(context.TODO(), newJokeBson)
	if err != nil {
		err := utils.NewError(http.StatusBadRequest, "Unable to insert joke to db")
		return c.JSON(http.StatusBadRequest, err)
	}

	// return inserted joke with new id
	var response model.Response
	response.Joke = newJoke
	response.ResponseMsg = "Success"
	response.StatusCode = http.StatusOK

	// return model.Response struct with inserted joke from request
	return c.JSON(http.StatusOK, response)
}
