package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "safepassword"
	DB_NAME     = "postgres"
	DB_HOST     = "localhost"
	DB_PORT     = "5430"
)

func ConnectToDb() *sql.DB {
	// connectionConfig := "user=postgres dbname=postgres password=safepassword host=localhost port=5430 sslmode=disable"
	connectionConfig := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT)
	db, err := sql.Open("postgres", connectionConfig)
	if err != nil {
		panic(err.Error())
	}
	return db
}
