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

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Load() Config {
	// Try to load .env file but don't fail if it doesn't exist
	_ = godotenv.Load()

	// Get required environment variables
	dbDNS := os.Getenv("DB_DNS")
	if dbDNS == "" {
		log.Fatal("DB_DNS environment variable is required")
	}

	// Get SMTP settings with defaults
	smtpHost := getEnv("SMTP_HOST", "smtp.gmail.com")
	smtpPort := getEnv("SMTP_PORT", "587")
	emailSender := getEnv("EMAIL_SENDER", "")
	emailPassword := getEnv("EMAIL_PASSWORD", "")

	// Validate email settings if they're not using defaults
	if emailSender == "" || emailPassword == "" {
		log.Fatal("EMAIL_SENDER and EMAIL_PASSWORD environment variables are required")
	}

	return Config{
		DBNS:          dbDNS,
		SMPTHost:      smtpHost,
		SMTPPort:      smtpPort,
		EMAILSender:   emailSender,
		EMAILPassword: emailPassword,
	}
}
