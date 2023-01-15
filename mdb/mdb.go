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
			if sqlError.Code != 1 {
				log.Fatal(sqlError)
			}
		} else {
			log.Fatal(err)
		}
	}
}

func emailEntryFromRow(row *sql.Rows) (*EmailEntry, error) {
	var id int64
	var email string
	var confirmedAt int64
	var optOut bool

	err := row.Scan(&id, &email, &confirmedAt, &optOut)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	t := time.Unix(confirmedAt, 0)
	return &EmailEntry{Id: id, Email: email, ConfirmAt: &t, OptOut: optOut}, nil
}

func CreateEmail(db *sql.DB, email string) error {
	_, err := db.Exec(`
	emails(email,confirmed_at,opt_out)
	VALUES(?,0,false)
	
	`, email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
