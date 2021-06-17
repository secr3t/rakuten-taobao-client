package client

import (
	"github.com/secr3t/rakuten-taobao-client/model"
	"strconv"
	"sync"
)

const defaultLimit = 400

type CompoundClient struct {
	ApiKey      string
	SearchLimit int
}

func NewCompoundClient(apiKey string, searchLimit int) *CompoundClient {
	if searchLimit > defaultLimit {
		searchLimit = defaultLimit
	}
	return &CompoundClient{
		ApiKey:      apiKey,
		SearchLimit: searchLimit,
	}
}

func (c *CompoundClient) SearchAndGetDetailsOneReqOneTime(param *SearchParam) chan model.Detail {
	var wg sync.WaitGroup
	limit := c.SearchLimit
	sc := NewSearchClient(c.ApiKey)

	search := sc.SearchItems(*param)

	detailChan := make(chan model.Detail, limit)

	wg.Add(1)
	go func() {
		if search.IsSuccess() {
			c.backgroundDetailRequestItemsSequential(&wg, search.Result.Item, detailChan)
			if limit > search.Result.TotalResults {
				limit = search.Result.TotalResults
			}

			pageSize, _ := strconv.Atoi(search.Result.PageSize)
			limit -= pageSize

			for ; limit > 0; limit -= pageSize {
				if param.Page == 0 {
					param.Page = 2
				} else {
					param.Page += 1
				}
				nextSearch := sc.SearchItems(*param)
				c.backgroundDetailRequestItemsSequential(&wg, nextSearch.Result.Item, detailChan)
			}
		}
		wg.Done()
	}()

	defer func() {
		go func() {
			wg.Wait()
			close(detailChan)
		}()
	}()

	return detailChan
}

func (c *CompoundClient) SearchAndGetDetailsMultiRequestOneTime(param *SearchParam) chan model.Detail {
	var wg sync.WaitGroup
	limit := c.SearchLimit
	sc := NewSearchClient(c.ApiKey)

	search := sc.SearchItems(*param)

	detailChan := make(chan model.Detail, limit)

	if search.IsSuccess() {
		c.backgroundDetailRequestItems(&wg, search.Result.Item, detailChan)
		if limit > search.Result.TotalResults {
			limit = search.Result.TotalResults
		}

		pageSize, _ := strconv.Atoi(search.Result.PageSize)
		limit -= pageSize

		for ; limit > 0; limit -= pageSize {
			if param.Page == 0 {
				param.Page = 2
			} else {
				param.Page += 1
			}
			nextSearch := sc.SearchItems(*param)
			c.backgroundDetailRequestItems(&wg, nextSearch.Result.Item, detailChan)
		}
	}
	defer func() {
		go func() {
			wg.Wait()
			close(detailChan)
		}()
	}()

	return detailChan
}

func (c *CompoundClient) backgroundDetailRequestItemsSequential(wg *sync.WaitGroup, items []model.SearchItem, ch chan model.Detail) {
	for _, item := range items {
		c.backgroundDetailRequestItem(wg, item, ch)
	}
}

func (c *CompoundClient) backgroundDetailRequestItems(wg *sync.WaitGroup, items []model.SearchItem, ch chan model.Detail) {
	for _, item := range items {
		go c.backgroundDetailRequestItem(wg, item, ch)
	}
}

func (c *CompoundClient) backgroundDetailRequestItem(wg *sync.WaitGroup, item model.SearchItem, ch chan model.Detail) {
	wg.Add(1)

	dc := NewDetailClient(c.ApiKey)

	detail := dc.GetDetail(item.NumIid)

	if detail.IsSuccess() {
		ch <- detail
	}

	wg.Done()
}
