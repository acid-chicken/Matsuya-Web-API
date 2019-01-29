package utils

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv .envを読み込み
func LoadEnv() {
	err := godotenv.Load()
	if err != nil && flag.Lookup("test.v") == nil {
		log.Print("Error loading .env file")
	}
}

// GetMongoConnectionURL MongoDBのURLを返す
func GetMongoConnectionURL() string {
	port := os.Getenv("MONGODB_CONNECTION_URL")
	if port == "" {
		return "mongodb://localhost/matsuya"
	}
	return port
}
