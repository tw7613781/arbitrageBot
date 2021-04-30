package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tw7613781/abitrage_bot/httpClient"
	"github.com/tw7613781/abitrage_bot/util"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_KEY := os.Getenv("API_KEY")
	API_SECRET := os.Getenv("API_SECRET")
	if API_KEY == "" || API_SECRET == "" {
		log.Fatal("API_KEY or API_SECRET is not provided")
	}

	config := util.GetConfig()

	c := httpClient.InitClient(API_KEY, API_SECRET, config.BaseURL)
	c.GetBalance("BTC")
}
