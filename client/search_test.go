package client

import "testing"

const apiKey = ""

func TestSearchClient_SearchItems(t *testing.T) {
	c := NewSearchClient(apiKey)

	p := NewSearchParam("one piece", "", 40, 0, 0, 0)

	search := c.SearchItems(p)
	t.Log(search)
}
