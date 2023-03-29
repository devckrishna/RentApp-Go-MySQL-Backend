package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var db *sql.DB

func GetDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env: ", err)
	}
	createdDatabase, err := sql.Open("mysql", os.Getenv("DSN"))

	if err != nil {
		log.Fatal("Failed to open DB Connection", err)
	}
	db = createdDatabase
	return db
}
