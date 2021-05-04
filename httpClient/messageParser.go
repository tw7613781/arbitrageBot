package httpClient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type Order struct {
	Quantity float64 `json:"Quantity"`
	Rate     float64 `json:"Rate"`
}

type OrderResult struct {
	Buy  []Order `json:"buy"`
	Sell []Order `json:"sell"`
}

type OrderBookBoth struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  OrderResult `json:"result"`
}

type OrderBookBuyOrSell struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Result  []Order `json:"result"`
}

type MarketResult struct {
	MarketCurrency     string `json:"MarketCurrency"`
	BaseCurrency       string `json:"BaseCurrency"`
	MarketCurrencyLong string `json:"MarketCurrencyLong"`
	BaseCurrencyLong   string `json:"BaseCurrencyLong"`
	MinTradeSize       string `json:"MinTradeSize"`
	MarketName         string `json:"MarketName"`
	IsActive           bool   `json:"IsActive"`
	Created            string `json:"Created"`
}

type Market struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Result  []MarketResult `json:"result"`
}

type TickerResult struct {
	Bid  float64 `json:"Bit"`
	Ask  float64 `json:"Ask"`
	Last float64 `json:"Last"`
}

type Ticker struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Result  TickerResult `json:"result"`
}

type BalanceResult struct {
	Currency      string  `json:"Currency"`
	Balance       float64 `json:"Balance"`
	Available     float64 `json:"Available"`
	Pending       float64 `json:"Pending"`
	CryptoAddress string  `json:"CryptoAddress"`
	Requested     bool    `json:"Requested"`
	Uuid          string  `json:"Uuid"`
}

type Balance struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Result  BalanceResult `json:"result"`
}

type LimitOrderResult struct {
	Uuid int `json:"uuid"`
}

type LimitOrder struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Result  LimitOrderResult `json:"result"`
}

func HttpRespToStruct(resp *http.Response, output interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Error with status code : %d", resp.StatusCode)
		return errors.New(string(resp.Status))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read req body error: %s", err)
	}

	err = json.Unmarshal(bodyBytes, output)
	if err != nil {
		log.Fatalf("Unmarshal data error: %s", err)
	}

	return nil
}
