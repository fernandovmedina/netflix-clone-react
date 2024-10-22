package database

/*
	-- date:   October 21, 23:13 (UTC-5)
	-- author: Fernando Vazquez
*/

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Function to connect the background's service with the mysql database
func ConnDB() (*sql.DB, error) {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		return nil, err
	}

	var (
		databaseName string = os.Getenv("DATABASE_NAME")
		databasePass string = os.Getenv("DATABASE_PASS")
		databaseUser string = os.Getenv("DATABASE_USER")
		databaseHost string = os.Getenv("DATABASE_HOST")
		databasePort string = os.Getenv("DATABASE_PORT")
	)

	var DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", databaseUser, databasePass, databaseHost, databasePort, databaseName)

	if DB, err = sql.Open("mysql", DSN); err != nil {
		return nil, err
	} else {
		if err = DB.Ping(); err != nil {
			return nil, err
		} else {
			return DB, nil
		}
	}
}
