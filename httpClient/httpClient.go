package httpClient

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"io/ioutil"
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

/*
* currency should be cryptocurrency symbol string. like "BTC", "ETH"
 */
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

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Read req body error: %s", err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
}

/*
* name can be any string,
* t should be 0 - general wallet, 1 - trade wallet
 */
func (c *client) AddWallet(name string, t uint8) {
	method := "/account/addwallet"

	params := &struct {
		Apikey string `url:"apikey"`
		Name   string `url:"name"`
		Nonce  string `url:"nonce"`
		Type   uint8  `url:"type"`
	}{
		c.apiKey,
		name,
		util.GetTimestampMili(),
		t,
	}

	resp, err := c.get(method, params)
	if err != nil {
		log.Fatalf("Add wallet error: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Read req body error: %s", err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
}

/*
* method -> "/markert/buylimit"
* params => "{market: 'dash-btc', quantity: '1', rate: '1'}" params except apikey and nonce and the params should strictly listed by alpha orders
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
	h.Write([]byte(c.baseURL + c.queryString))
	sha := hex.EncodeToString(h.Sum(nil))

	req.Header.Add("apisign", sha)
	log.Println(req.URL)
	return c.c.Do(req)
}
