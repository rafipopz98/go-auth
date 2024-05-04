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
	connStr := "postgres://avnadmin:AVNS_0OWTMktf96_v4t_iPEg@midas-lab-project-auth.c.aivencloud.com:16898/defaultdb?sslmode=require"
	// Connect to database
	database, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.db = database
	fmt.Println("Connected successfully")
	return nil
}
