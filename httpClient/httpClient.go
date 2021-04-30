package httpClient

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"

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
		Apikey   string `json:"apikey"`
		Currency string `json:"currency"`
		Nonce    string `json:"nonce"`
	}{
		Apikey:   c.apiKey,
		Currency: currency,
		Nonce:    util.GetTimestampMili(),
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
	paramsMarshal, err := json.Marshal(params)

	if err != nil {
		log.Fatalf("Fial to parse params: %s", err)
	}

	u, err := url.Parse(c.baseURL)
	if err != nil {
		log.Fatalf("Fail to parse base url: %s", err)
	}

	urlRemain := method + "?" + string(paramsMarshal)
	u.Path = path.Join(u.Path, urlRemain)

	c.queryString = u.String()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return c.do(req)
}

func (c *client) do(req *http.Request) (*http.Response, error) {
	h := hmac.New(sha512.New, []byte(c.apiSecret))
	h.Write([]byte(c.queryString))
	sha := hex.EncodeToString(h.Sum(nil))

	req.Header.Add("apisign", sha)
	return c.c.Do(req)
}
