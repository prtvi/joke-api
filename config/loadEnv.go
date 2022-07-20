package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// load env variables into the system env
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}
}