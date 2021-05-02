package httpClient_test

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/tw7613781/arbitrageBot/httpClient"
	"github.com/tw7613781/arbitrageBot/util"
)

func initConfig(t *testing.T, API_KEY *string, API_SECRET *string) *util.Config {
	err := godotenv.Load("./../.env")
	assert.Nil(t, err, "Error loading .env file")

	*API_KEY = os.Getenv("API_KEY")
	*API_SECRET = os.Getenv("API_SECRET")
	assert.NotEqual(t, *API_KEY, "")
	assert.NotEqual(t, *API_SECRET, "")

	return util.GetConfig("./..")
}
func TestGetOrderBookBuyOrSell(t *testing.T) {
	var API_KEY, API_SECRET string
	config := initConfig(t, &API_KEY, &API_SECRET)
	c := httpClient.InitClient(API_KEY, API_SECRET, config.BaseURL)
	result, err := c.GetOrderBookBuyOrSell("krw-eth", "buy")
	log.Println(result)
	assert.Nil(t, err)
}

func TestGetOrderBoth(t *testing.T) {
	var API_KEY, API_SECRET string
	config := initConfig(t, &API_KEY, &API_SECRET)
	c := httpClient.InitClient(API_KEY, API_SECRET, config.BaseURL)
	result, err := c.GetOrderBookBoth("krw-eth")
	log.Println(result)
	assert.Nil(t, err)
}

func TestGetTicker(t *testing.T) {
	var API_KEY, API_SECRET string
	config := initConfig(t, &API_KEY, &API_SECRET)
	c := httpClient.InitClient(API_KEY, API_SECRET, config.BaseURL)
	result, err := c.GetTicker("krw-eth")
	log.Println(result)
	assert.Nil(t, err)
}

func TestGetBalance(t *testing.T) {
	var API_KEY, API_SECRET string
	config := initConfig(t, &API_KEY, &API_SECRET)
	c := httpClient.InitClient(API_KEY, API_SECRET, config.BaseURL)
	result, err := c.GetBalance("BTC")
	log.Println(result)
	assert.Nil(t, err)
}

func TestGetMarkets(t *testing.T) {
	var API_KEY, API_SECRET string
	config := initConfig(t, &API_KEY, &API_SECRET)
	c := httpClient.InitClient(API_KEY, API_SECRET, config.BaseURL)
	result, err := c.GetMarkets()
	log.Println(result)
	assert.Nil(t, err)
}
