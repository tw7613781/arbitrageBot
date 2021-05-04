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
	// 0.1% trading feee
	fee := 0.001
	for {
		balance, err := c.GetBalance("KRW")
		if err != nil {
			log.Printf("error get balance of KRW: %v", err)
		}
		log.Println("--------------------------------------")
		log.Printf("krw balance: %v", balance.Available)

		tmp := balance.Available
		quanlityNotEnough := false

		for _, pair := range pairs {
			r, err := c.GetOrderBookBuyOrSell(pair, "buy")
			if err != nil {
				log.Printf("error to get order book sell price of %v: %v", pair, err)
			}
			// log.Println(r)
			log.Printf("pair %v price - %v, quanlity - %v", pair, r[0].Rate, r[0].Quantity)
			if tmp <= r[0].Quantity {
				tmp = tmp * r[0].Rate * (1 - fee)
				log.Printf("tmp value: %v", tmp)
			} else {
				log.Printf("pair %v: the quantily is %v, less then requested %v", pair, r[0].Quantity, tmp)
				quanlityNotEnough = true
				break
			}
		}
		if quanlityNotEnough {
			continue
		}
		// log.Printf("New balance: %v", tmp)
		// log.Printf("Old balance: %v", balance.Available)
		// log.Printf("Change rate: %v", (tmp-balance.Balance)/balance.Available*100)
		changeRate := (tmp - balance.Balance) / balance.Available
		if changeRate > 0 {
			log.Printf("New balance: %v", tmp)
			log.Printf("Old balance: %v", balance.Available)
			log.Printf("Found chance: %v", changeRate)
		}
		log.Printf("Change Rate: %v", changeRate)
		time.Sleep(3 * time.Second)
	}

}
