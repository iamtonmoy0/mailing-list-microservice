package mdb

import (
	"database/sql"
	"log"
	"time"

	"github.com/mattn/go-sqlite3"
)

type EmailEntry struct {
	Id        int64
	Email     string
	ConfirmAt *time.Time
	OptOut    bool
}

func TryCreate(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE emails (
		id INTEGER PRIMARY KEY,
		email TEXT UNIQUE,
		confirm_at INTEGER,
		opt_opt INTEGER
	);
	`)
	if err != nil {
		if sqlError, ok := err.(sqlite3.Error); ok {
			//code 1 == "table already exist"
			if sqlError.code != 1 {
				log.Fatal(sqlError)
			}
		}els
	}
}
