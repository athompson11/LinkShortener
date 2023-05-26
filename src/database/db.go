package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DBConn *sql.DB

func InitDatabase(databaseLocation string) {
	var err error
	DBConn, err = sql.Open("sqlite3", databaseLocation)
	if err != nil {
		log.Fatal(err)
	}
}
