package client

import (
	"encoding/json"
	"github.com/secr3t/rakuten-taobao-client/model"
	"io/ioutil"
	"net/http"
	"sync"
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

func (c *SearchClient) SearchTilLimit(param *SearchParam, limit int) []model.SearchItem {
	result := c.SearchItems(*param)

	if limit > result.Result.TotalResults {
		limit = result.Result.TotalResults
	}

	itemsChain := make(chan model.SearchItem, limit)

	limit -= param.PageSize

	for _, item := range result.Result.Item {
		itemsChain <- item
	}

	wg := sync.WaitGroup{}

	for ;limit >0; limit -= param.PageSize {
		param.Page += 1
		wg.Add(1)
		go func() {
			result = c.SearchItems(*param)
			for _, item := range result.Result.Item {
				itemsChain <- item
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(itemsChain)
	}()

	items := make([]model.SearchItem, 0)
	for item := range itemsChain {
		items = append(items, item)
	}

	return items
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

	var search model.Search

	json.Unmarshal(body, &search)

	rateLimit := model.FromHeader(res.Header)
	search.RateLimit = rateLimit

	return search
}
