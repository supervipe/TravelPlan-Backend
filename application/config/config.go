package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	db "backend/infra/db"
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

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	var dbType string
	if len(params) > 0 {
		dbType = params[0]
	} else {
		dbType = db.GetDBType()
	}
	if dbType == "postgres" {
		DB, err = gorm.Open(postgres.Open(db.GetPostgresConnectionString()), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}
		log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile).Println("Connected to Postgres")
	}
	return DB
}
