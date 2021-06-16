package client

import "testing"

func TestDetailClient_GetDetail(t *testing.T) {
	c := NewDetailClient(apiKey)

	detail := c.GetDetail(647673635989)

	t.Log(detail)
}
