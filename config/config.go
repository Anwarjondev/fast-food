package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBNS          string
	SMPTHost      string
	SMTPPort      string
	EMAILSender   string
	EMAILPassword string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		DBNS:          os.Getenv("DB_DNS"),
		SMPTHost:      os.Getenv("SMTP_HOST"),
		SMTPPort:      os.Getenv("SMTP_PORT"),
		EMAILSender:   os.Getenv("EMAIL_SENDER"),
		EMAILPassword: os.Getenv("EMAIL_PASSWORD"),
	}
}
