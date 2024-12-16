package bd

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConDB() (*sqlx.DB, error) {
	godotenv.Load()

	dsn := os.Getenv("DSN")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}
	return db, nil
}
