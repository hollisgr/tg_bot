package cfg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetToken() string {
	err := godotenv.Load("cfg.env")
	if err != nil {
		log.Fatal("cfg.env not loaded")
	}

	return os.Getenv("TELEGRAM_BOT_TOKEN")
}

func GetURL() string {
	err := godotenv.Load("cfg.env")
	if err != nil {
		log.Fatal("cfg.env not loaded")
	}

	return os.Getenv("API_URL")
}
