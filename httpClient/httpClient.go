package httpClient

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/tw7613781/abitrage_bot/util"
)

type client struct {
	c           http.Client
	baseURL     string
	apiKey      string
	apiSecret   string
	queryString string
}

func InitClient(apiKey string, apiSecret string, baseURL string) *client {
	c := client{
		c:         http.Client{},
		apiKey:    apiKey,
		apiSecret: apiSecret,
		baseURL:   baseURL,
	}
	return &c
}

func (c *client) GetBalance(currency string) {
	method := "/account/getbalance"

	params := &struct {
		Apikey   string `url:"apikey"`
		Currency string `url:"currency"`
		Nonce    string `url:"nonce"`
	}{
		c.apiKey,
		currency,
		util.GetTimestampMili(),
	}

	resp, err := c.get(method, params)
	if err != nil {
		log.Fatalf("Get balance error: %s", err)
	}
	fmt.Println(resp)
}

/*
* method -> "/markert/buylimit"
* params => "{market: 'dash-btc', quantity: '1', rate: '1'}" params except apikey and nonce
 */
func (c *client) get(method string, params interface{}) (resp *http.Response, err error) {
	v, err := query.Values(params)
	if err != nil {
		log.Fatalf("Fail to parse params: %s", err)
	}

	c.queryString = method + "?" + v.Encode()

	req, err := http.NewRequest("GET", c.baseURL+c.queryString, nil)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *client) do(req *http.Request) (*http.Response, error) {
	h := hmac.New(sha512.New, []byte(c.apiSecret))
	h.Write([]byte("https://api.dovewallet.com/v1.1" + c.queryString))
	sha := hex.EncodeToString(h.Sum(nil))

	req.Header.Add("apisign", sha)
	fmt.Println(req)
	return c.c.Do(req)
}
