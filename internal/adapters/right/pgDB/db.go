package pgdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func NewDB() *DB {
	return &DB{}
}

func (db *DB) MakeDBConnection() error {
	connStr := "kk"
	// Connect to database
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.db = database
	fmt.Println("Connected successfully")
	return nil
}
