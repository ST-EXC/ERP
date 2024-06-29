package data

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB(db *sql.DB) {
	DB = db
	log.Println("Database initialized")
}
