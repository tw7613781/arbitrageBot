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

	var c httpClient.Client
	c.ApiKey = API_KEY
	c.ApiSecret = API_SECRET
	c.Base_url = config.BaseURL

	method := "/account/getbalance"
	param := map[string]interface{}{
		"currency": "btc",
	}

	c.Get(method, param)

}
