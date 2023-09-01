package config

import (
	"os"

	"github.com/joho/godotenv"
)

func ConfEnv() {
	mode := os.Args[1]
	if mode == "dev" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
		return
	}

	if mode == "prod" {
		err := godotenv.Load(".env_prod")
		if err != nil {
			panic(err)
		}
		return
	}

	panic("No environment set")
}

func GetPort() string {
	return os.Getenv("PORT")
}
