package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnMongoDB() (*mongo.Client, error) {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		return nil, err
	}

	var (
		databaseUser string = os.Getenv("DATABASE_USER")
		databasePass string = os.Getenv("DATABASE_PASS")
		databaseHost string = os.Getenv("DATABASE_HOST")
		databasePort string = os.Getenv("DATABASE_PORT")
	)

	var dsn = fmt.Sprintf("mongodb://%s:%s@%s:%s", databaseUser, databasePass, databaseHost, databasePort)

	var clientOptions = options.Client().ApplyURI(dsn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %s", err.Error())
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("couldn't ping MongoDB: %s", err.Error())
	}

	fmt.Printf("Successfully connected to MongoDB: %s\n", client.Database("netflix_clone").Name())

	return client, nil
}
