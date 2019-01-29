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
	url := os.Getenv("MONGODB_CONNECTION_URL")
	if url == "" {
		return "mongodb://localhost/matsuya"
	}
	return url
}

// GetDBName MongoDBのDB名を返す
func GetDBName() string {
	name := os.Getenv("MONGODB_DATABASE_NAME")
	if name == "" {
		return "matsuya"
	}
	return name
}
