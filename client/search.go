package client

import (
	"encoding/json"
	"fmt"
	"github.com/secr3t/rakuten-taobao-client/model"
	"io/ioutil"
	"net/http"
)

const (
	taobaoApiHost = "taobao-api.p.rapidapi.com"
	searchItems   = "item_search"
)

type SearchClient struct {
	ApiKey string
}

func NewSearchClient(apiKey string) *SearchClient {
	return &SearchClient{
		ApiKey: apiKey,
	}
}

func (c *SearchClient) SearchItems(param SearchParam) model.Search {
	query := param.ToQuery()

	uri := GetUri(taobaoApiHost, searchItems, query)

	req, _ := http.NewRequest("GET", uri, nil)

	req.Header.Add("x-rapidapi-key", c.ApiKey)
	req.Header.Add("x-rapidapi-host", taobaoApiHost)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	var search model.Search

	json.Unmarshal(body, &search)

	return search
}
