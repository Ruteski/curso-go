package main

import (
	"auction_go/configuration/database/mongodb"
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error trying to load .env file")
		return
	}

	databaseClient, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
