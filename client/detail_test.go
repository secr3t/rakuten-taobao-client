package client

import "testing"

func TestDetailClient_GetDetail(t *testing.T) {
	c := NewDetailClient(apiKey)

	detail := c.GetDetail(613371439320)

	t.Log(detail)
}
