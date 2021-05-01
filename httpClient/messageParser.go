package httpClient

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type Order struct {
	Quantity float32 `json:"Quantity"`
	Rate     float32 `json:"Rate"`
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
	Bid  float32 `json:"Bit"`
	Ask  float32 `json:"Ask"`
	Last float32 `json:"Last"`
}

type Ticker struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Result  TickerResult `json:"result"`
}

type BalanceResult struct {
	Currency      string  `json:"Currency"`
	Balance       float32 `json:"Balance"`
	Available     float32 `json:"Available"`
	Pending       float32 `json:"Pending"`
	CryptoAddress string  `json:"CryptoAddress"`
	Requested     bool    `json:"Requested"`
	Uuid          string  `json:"Uuid"`
}

type Balance struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Result  BalanceResult `json:"result"`
}

func HttpRespToOrderBookBoth(resp *http.Response) (OrderResult, error) {
	defer resp.Body.Close()

	var d OrderBookBoth
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error with status code : %d", resp.StatusCode)
		return d.Result, errors.New(string(resp.Status))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read req body error: %s", err)
	}

	err = json.Unmarshal(bodyBytes, &d)
	if err != nil {
		log.Fatalf("Unmarshal data error: %s", err)
	}
	if !d.Success {
		log.Printf("Error with data message: %s", d.Message)
		return d.Result, errors.New(d.Message)
	}
	return d.Result, nil
}

func HttpRespToOrderBuyOrSell(resp *http.Response) ([]Order, error) {
	defer resp.Body.Close()

	var d OrderBookBuyOrSell
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error with status code : %d", resp.StatusCode)
		return d.Result, errors.New(string(resp.Status))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read req body error: %s", err)
	}

	err = json.Unmarshal(bodyBytes, &d)
	if err != nil {
		log.Fatalf("Unmarshal data error: %s", err)
	}
	if !d.Success {
		log.Printf("Error with data message: %s", d.Message)
		return d.Result, errors.New(d.Message)
	}
	return d.Result, nil
}

func HttpRespToMarket(resp *http.Response) ([]MarketResult, error) {
	defer resp.Body.Close()

	var d Market
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error with status code : %d", resp.StatusCode)
		return d.Result, errors.New(string(resp.Status))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read req body error: %s", err)
	}

	err = json.Unmarshal(bodyBytes, &d)
	if err != nil {
		log.Fatalf("Unmarshal data error: %s", err)
	}
	if !d.Success {
		log.Printf("Error with data message: %s", d.Message)
		return d.Result, errors.New(d.Message)
	}
	return d.Result, nil
}

func HttpRespToTicker(resp *http.Response) (TickerResult, error) {
	defer resp.Body.Close()

	var d Ticker
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error with status code : %d", resp.StatusCode)
		return d.Result, errors.New(string(resp.Status))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read req body error: %s", err)
	}

	err = json.Unmarshal(bodyBytes, &d)
	if err != nil {
		log.Fatalf("Unmarshal data error: %s", err)
	}
	if !d.Success {
		log.Printf("Error with data message: %s", d.Message)
		return d.Result, errors.New(d.Message)
	}
	return d.Result, nil
}

func HttpRespToBalance(resp *http.Response) (BalanceResult, error) {
	defer resp.Body.Close()

	var d Balance
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error with status code : %d", resp.StatusCode)
		return d.Result, errors.New(string(resp.Status))
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read req body error: %s", err)
	}

	err = json.Unmarshal(bodyBytes, &d)
	if err != nil {
		log.Fatalf("Unmarshal data error: %s", err)
	}
	if !d.Success {
		log.Printf("Error with data message: %s", d.Message)
		return d.Result, errors.New(d.Message)
	}
	return d.Result, nil
}
