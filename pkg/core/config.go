package core

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DBURL string
}

func LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Port not found in environment")
	}

	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("DB URL not found in environment")
	}

	return Config{
		Port:  portString,
		DBURL: dbString,
	}
}
