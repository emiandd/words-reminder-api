package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)


func DBConnector() (*sql.DB, error) {

	var db *sql.DB
	cfg := mysql.Config{
		User:   os.Getenv("USER_DB"),
		Passwd: os.Getenv("PASS_DB"),
		Net:    os.Getenv("NET"),
		Addr:   os.Getenv("HOST"),
		DBName: os.Getenv("DBNAME"),
	}

	// Get a driver-specific connector.
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		log.Fatal(err)
		return db, err
	}

	// Get a database handle.
	db = sql.OpenDB(connector)

	return db, nil
}
