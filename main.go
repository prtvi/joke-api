package main

import (
	"os"

	config "prtvi/joke-api/config"
	routes "prtvi/joke-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// load the environment variables
	config.LoadEnv()

	PORT := os.Getenv("PORT")

	config.EstablishConnection()

	e := echo.New()

	e.GET("/", routes.Welcome)

	// READ
	e.GET("/random", routes.RandomOne)
	e.GET("/random/:n", routes.RandomN)
	e.GET("/joke/:id", routes.GetByID)

	// CREATE
	e.POST("/new", routes.AddOne)

	// DELETE
	e.DELETE("/remove/:id", routes.DeleteOne)

	// UPDATE
	e.PUT("/update/:id", routes.UpdateOne)

	e.Logger.Fatal(e.Start(":" + PORT))
}
