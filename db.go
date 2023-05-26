package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DBConn *sql.DB

func InitDatabase() {
	var err error
	DBConn, err = sql.Open("sqlite3", os.Getenv("DATABASE_LOCATION"))
	if err != nil {
		log.Fatal(err)
	}
}
