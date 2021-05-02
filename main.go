package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tw7613781/arbitrageBot/arbitrage"
	"github.com/tw7613781/arbitrageBot/httpClient"
	"github.com/tw7613781/arbitrageBot/util"
)

func initConfig(API_KEY *string, API_SECRET *string) *util.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error to setup env: %v", err)
	}
	*API_KEY = os.Getenv("API_KEY")
	*API_SECRET = os.Getenv("API_SECRET")

	return util.GetConfig("./")
}

func main() {
	var API_KEY, API_SECRET string
	config := initConfig(&API_KEY, &API_SECRET)
	c := httpClient.InitClient(API_KEY, API_SECRET, config.BaseURL)

	pairs := []string{"eth-krw", "btc-eth", "krw-btc"}
	arbitrage.FindChance(c, pairs)
}
