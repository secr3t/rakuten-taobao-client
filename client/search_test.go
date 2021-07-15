package client

import (
	"github.com/secr3t/rakuten-taobao-client/model"
	"testing"
)

const apiKey = ""

func TestSearchClient_SearchItems(t *testing.T) {
	c := NewSearchClient(apiKey)

	p := NewSearchParam("鹦鹉洗澡盆盒", "", 0, 0, 0, 0, 0)

	t.Log(model.ToJson(p))

	search := c.SearchItems(p)
	t.Log(model.ToJson(search))
}
