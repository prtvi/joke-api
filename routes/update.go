package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	config "prtvi/joke-api/config"
	model "prtvi/joke-api/model"
	utils "prtvi/joke-api/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// "/update/:id"
// update entire joke with request body

func UpdateOne(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("Endpoint hit: /update/%v\n", id)

	// read request body
	var jokeToUp model.Joke
	err := json.NewDecoder(c.Request().Body).Decode(&jokeToUp)

	// if request body is empty
	if utils.IfBodyContent(jokeToUp, err) {
		err := utils.NewError(http.StatusBadRequest, "No request body found")
		return c.JSON(http.StatusBadRequest, err)
	}

	// create filter, update and options for querying
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.D{
			{Key: "id", Value: id},
			{Key: "type", Value: jokeToUp.Type},
			{Key: "punchline", Value: jokeToUp.Punchline},
			{Key: "setup", Value: jokeToUp.Setup}},
	}
	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	cursor := config.Collection.FindOneAndUpdate(context.TODO(), filter, update, &opt)
	if cursor.Err() != nil {
		err := utils.NewError(http.StatusBadRequest, "Could not find joke to update")
		return c.JSON(http.StatusBadRequest, err)
	}

	// decode fetched doc
	doc := bson.M{}
	decodeErr := cursor.Decode(&doc)
	if decodeErr != nil {
		err := utils.NewError(http.StatusBadRequest, "Could not find joke to update")
		return c.JSON(http.StatusBadRequest, err)
	}

	// response obj to send back updated joke
	var response model.Response
	response.Joke = utils.MongoDocToStruct(doc)
	response.StatusCode = http.StatusOK
	response.ResponseMsg = "Success"

	return c.JSON(http.StatusOK, response)
}
