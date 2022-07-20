package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	config "prtvi/joke-api/config"
	utils "prtvi/joke-api/utils"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

// "/delete/:id"
// delete joke by id

func DeleteOne(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("Endpoint hit: /delete/%v\n", id)

	cursor := config.Collection.FindOneAndDelete(context.TODO(), bson.D{{Key: "id", Value: id}})
	if cursor.Err() != nil {
		err := utils.NewError(http.StatusBadRequest, "No joke with the given id exists!")
		return c.JSON(http.StatusBadRequest, err)
	}

	success := utils.NewError(http.StatusOK, "Delete operation successful!")
	return c.JSON(http.StatusOK, success)
}
