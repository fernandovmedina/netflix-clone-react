package database

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnDB() (*sql.DB, error) {
	var err error

	if err = godotenv.Load(); err != nil {
		return nil, err
	}

	var (
		databaseName string = os.Getenv("DATABASE_NAME")
		databaseUser string = os.Getenv("DATABASE_USER")
		databasePass string = os.Getenv("DATABASE_PASS")
		databasePort string = os.Getenv("DATABASE_PORT")
	)

	var dsn string = databaseUser + ":" + databasePass + "@tcp(127.0.0.1:" + databasePort + ")/" + databaseName

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
