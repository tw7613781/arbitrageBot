package arbitrage

import (
	"log"
	"time"

	"github.com/tw7613781/arbitrageBot/httpClient"
)

/*
* c is the custimised http client
* pairs are the 3 trading pairs for triangle arbitrage. like ["krw-usdt", "usdt-eth", "eth-krw"]
* but for the dove wallet order book naming convesion of pair, we need to give the pair order like ["usdt-krw", "eth-usdt", "krw-eth"]
* the fee is 0.1%
 */
func FindChance(c *httpClient.Client, pairs []string) {
	for {
		balance, err := c.GetBalance("KRW")
		if err != nil {
			log.Printf("error get balance of KRW: %v", err)
		}
		log.Printf("krw balance: %v", balance.Available)
		for i, pair := range pairs {
			r, err := c.GetOrderBookBuyOrSell(pair, "sell")
			if err != nil {
				log.Printf("%v: error to get order book sell price of %v: %v", i, pair, err)
			}
			log.Printf("pair %v price - %v, quanlity - %v", pair, r[0].Rate, r[0].Quantity)
		}
		time.Sleep(3 * time.Second)
	}

}
