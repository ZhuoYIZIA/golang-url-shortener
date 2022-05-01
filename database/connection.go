package database

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"os"
)

func PostgreSQLConnect() (*sql.DB, error) {
	dsn := getPostgreSQLDsn()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		// TODO: error handler
		return nil, err
	}

	return db, nil
}

func getPostgreSQLDsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Taipei",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PASSWORD"),
	)
}
