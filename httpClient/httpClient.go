package httpClient

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/tw7613781/abitrage_bot/util"
)

type Client struct {
	c           http.Client
	Base_url    string
	ApiKey      string
	ApiSecret   string
	queryString string
}

/*
* method -> "/markert/buylimit"
* params => "{market: 'dash-btc', quantity: '1', rate: '1'}" params except apikey and nonce
 */
func (c *Client) Get(method string, params map[string]interface{}) (resp *http.Response, err error) {
	params["apikey"] = c.ApiKey
	params["nonce"] = time.Now().UnixNano() / int64(time.Millisecond)

	paramsSerialize := util.SortMapByKeyToString(params)

	u, err := url.Parse(c.Base_url)
	if err != nil {
		log.Fatalf("Fail to parse base url: %s", err)
	}

	urlRemain := method + "?" + string(paramsSerialize)
	u.Path = path.Join(u.Path, urlRemain)

	c.queryString = u.String()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return c.Do(req)
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	h := hmac.New(sha512.New, []byte(c.ApiSecret))
	h.Write([]byte(c.queryString))
	sha := hex.EncodeToString(h.Sum(nil))

	req.Header.Add("apisign", sha)
	return c.c.Do(req)
}
