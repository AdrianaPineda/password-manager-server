package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func InitDB(dataSourceName string) (*sql.DB, error) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db, err
}
