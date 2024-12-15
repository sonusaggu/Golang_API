package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sqlx.DB, error) {
	//connStr := os.Getenv("DB_CONNECTION_STRING")
	connStr := "user=postgres password=mysecretpassword dbname=postgres sslmode=disable"
	if connStr == "" {
		connStr = "user=postgres password=yourpassword dbname=yourdbname sslmode=disable"
	}
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
		return nil, err
	}
	return db, nil

}
