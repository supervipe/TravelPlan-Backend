package db

import "fmt"

const (
	DBUser     = "postgres"
	DBPassword = "postgres"
	DBName     = "postgres"
	DBHost     = "0.0.0.0"
	DBPort     = "5432"
	DBType     = "postgres"
)

func GetDBType() string {
	return DBType
}

func GetPostgresConnectionString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DBHost, DBUser, DBPassword, DBName, DBPort)
}
