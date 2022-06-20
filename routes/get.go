package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	config "jokeapi/config"
	utils "jokeapi/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// "/random" route
// returns a random joke

func RandomOne(c echo.Context) error {
	fmt.Println("Endpoint hit: /random")

	results := utils.GetJokes(1, config.Collection)
	return c.JSON(http.StatusOK, results[0])
}

// "/random/:n"
// returns n random jokes

func RandomN(c echo.Context) error {
	n, _ := strconv.Atoi(c.Param("n"))
	fmt.Printf("Endpoint hit: /random/%v\n", n)

	results := utils.GetJokes(n, config.Collection)
	return c.JSON(http.StatusOK, results)
}

// "/joke/:id"
// returns a joke by id

func GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("Endpoint hit: /joke/%v\n", id)

	// check if request id is less than # of documents in db
	if count, _ := config.Collection.CountDocuments(context.TODO(), bson.D{}); id <= 0 || id > int(count) {
		err := utils.NewError(http.StatusBadRequest, "No joke with this ID exists!")
		return c.JSON(http.StatusBadRequest, err)
	}

	// find by id
	cursor := config.Collection.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}})
	if cursor.Err() != nil {
		err := utils.NewError(http.StatusNotFound, "Joke not found")
		return c.JSON(http.StatusNotFound, err)
	}

	// decode the fetched document
	result := bson.M{}
	decodeErr := cursor.Decode(&result)
	if decodeErr != nil {
		err := utils.NewError(http.StatusNotFound, "Joke not found")
		return c.JSON(http.StatusNotFound, err)
	}

	// convert mongo doc to model.Joke struct
	return c.JSON(http.StatusOK, utils.MongoDocToStruct(result))
}
