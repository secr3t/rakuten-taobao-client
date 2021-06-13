package client

import (
	"encoding/json"
	"fmt"
	"github.com/secr3t/rakuten-taobao-client/model"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	taobaoAdvancedHost = "taobao-advanced.p.rapidapi.com"
	itemDetail         = "item_detail"
)

type DetailClient struct {
	ApiKey string
}

func NewDetailClient(apiKey string) *DetailClient {
	return &DetailClient{
		ApiKey: apiKey,
	}
}

func (c *DetailClient) GetDetail(numiid int64) model.Detail {
	query := fmt.Sprintf("num_iid=%d", numiid)
	uri := GetUri(taobaoAdvancedHost, itemDetail, query)

	req, _ := http.NewRequest("GET", uri, nil)

	req.Header.Add("x-rapidapi-key", c.ApiKey)
	req.Header.Add("x-rapidapi-host", taobaoAdvancedHost)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var detail model.Detail

	json.Unmarshal(body, &detail)

	if !detail.IsSuccess() {
		log.Printf("GetDetail Failed num_iid = %d, response = %s", numiid, string(body))
	}

	return detail
}
