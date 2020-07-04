package config

import (
	"os"

	"github.com/joho/godotenv"
)

func InitConfig() {
	env := os.Getenv("GO_ENV")

	if "" == env {
		env = "development"
	}

	err := godotenv.Load(".env." + env)
	if err != nil {
		panic(err.Error())
	}
}
