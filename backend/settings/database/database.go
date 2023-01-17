package database

import (
	"database/sql"

	"backend/settings"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// create the dsn string from variables that are specified in the .env file
func DsnString() string {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		settings.DB_HOST,
		settings.DB_USER,
		settings.DB_PASS,
		settings.DB_NAME,
		settings.DB_PORT,
		settings.DB_SSL,
		settings.DB_TZ)
	return dsn
}

// call a db connection without using the global db variable
func GetDbConn() (*sql.DB, error) {

	db, err := sql.Open("postgres", DsnString())

	if err != nil {
		panic(err)
	}
	return db, err

}

func Connect() {

	db, err := sql.Open("postgres", DsnString())

	if err != nil {
		panic(err)
	}

	// assign *sql.DB to the global variable
	DB = db

}
