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
	var fee, initialKrw float64 = 0.001, 10000
	for {
		log.Printf("%v ...", pairs)
		stopCurrentRound := false
		tmp := initialKrw
		var orders []httpClient.Order
		for _, pair := range pairs {
			r, err := c.GetOrderBookBuyOrSell(pair, "buy")
			if err != nil {
				log.Printf("error to get order book sell price of %v: %v", pair, err)
				stopCurrentRound = true
				break
			}
			// log.Printf("pair %v price - %v, quanlity - %v", pair, r[0].Rate, r[0].Quantity)
			if tmp <= r[0].Quantity {
				tmp = tmp * r[0].Rate * (1 - fee)
				order := httpClient.Order{
					Quantity: tmp,
					Rate:     r[0].Rate,
				}
				orders = append(orders, order)
				// log.Printf("tmp value: %v", tmp)
			} else {
				log.Printf("pair %v: the quantily is %v, less then requested %v", pair, r[0].Quantity, tmp)
				stopCurrentRound = true
				break
			}
		}
		if stopCurrentRound {
			continue
		}
		changeRate := (tmp - initialKrw) / initialKrw
		log.Printf("%v: %v", pairs, changeRate)
		if changeRate > 0 {
			log.Printf("New balance: %v", tmp)
			log.Printf("Old balance: %v", initialKrw)
			log.Printf("Found chance: %v", changeRate)
			execute(c, pairs, orders)
		}
		time.Sleep(3 * time.Second)
	}
}

func execute(c *httpClient.Client, pairs []string, orders []httpClient.Order) {
	for i := 0; i < 3; i++ {
		_, err := c.LimitOrder(pairs[0], orders[0].Rate, orders[0].Quantity, "sell")
		if err != nil {
			log.Printf("Order for %v failed: %v", pairs[0], err)
			break
		}
	}
	log.Printf("Orders %v succeed %v", pairs, orders)
}
