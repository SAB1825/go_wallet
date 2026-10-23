package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DBDSN string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Env file not loaded, using system environmental variables")
		return nil, err
	}

	dsn := os.Getenv("DB_USER`") + ":" +
		os.Getenv("DB_PASSWORD") + "@tcp(" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") + "?parseTime=true"

	config := &Config{
		DBDSN: dsn,
	}

	return config, nil
}
