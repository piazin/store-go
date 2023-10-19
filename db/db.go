package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/piazin/store-go/utils"
)

func ConnectToDatabase() *sql.DB {
	connectionStr := "user=postgres dbname=store_go password=postgres host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	utils.CheckError(err)

	return db
}