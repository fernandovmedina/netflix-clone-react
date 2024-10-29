package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// Function to connect to mysql database
func ConnMYSQL() (*sql.DB, error) {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		return nil, err
	}

	var (
		databaseUser string = os.Getenv("DATABASE_USER")
		databasePass string = os.Getenv("DATABASE_PASS")
		databaseName string = os.Getenv("DATABASE_NAME")
		databaseHost string = os.Getenv("DATABASE_HOST")
		databasePort string = os.Getenv("DATABASE_PORT")
	)

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", databaseUser, databasePass, databaseHost, databasePort, databaseName)

	if DB, err = sql.Open("mysql", dsn); err != nil {
		return nil, err
	} else {
		if err = DB.Ping(); err != nil {
			return nil, err
		} else {
			return DB, nil
		}
	}
}
