package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() *sql.DB {
	connectionStr := "user=postgres dbname=store_go password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)

	if err != nil {
		panic(err.Error())
	}

	return db
}