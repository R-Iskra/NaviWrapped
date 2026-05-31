package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	_ "embed"
)

//go:embed schema.sql
var schema string

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite", "./naviwrapped.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
