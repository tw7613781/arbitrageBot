package main

import (
	"fmt"

	"github.com/tw7613781/abitrage_bot/config"
	"github.com/tw7613781/abitrage_bot/httpclient"
)

func main() {
	API_KEY := config.GetEnv("API_KEY")

	fmt.Println("Hello dove wallet: ", API_KEY)
	httpclient.Get()
}
