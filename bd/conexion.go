package bd

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func ConDB() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	dsn := os.Getenv("DSN")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db, nil
}
