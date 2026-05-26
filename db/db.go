package db

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	_ "embed"
)

//go:embed schema.sql
var schema string

func Init() {
	db, err := sql.Open("sqlite", "./naviwrapped.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
