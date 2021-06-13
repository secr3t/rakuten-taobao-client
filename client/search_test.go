package client

import "testing"

const apiKey = "b4b579ca92mshe0e0b0bf9017014p15408djsnbcc197b35d80"

func TestSearchClient_SearchItems(t *testing.T) {
	c := NewSearchClient(apiKey)

	p := NewSearchParam("onepiece", "", 0, 0, 0)

	search := c.SearchItems(p)
	t.Log(search)
}
