package client

import (
	"fmt"
	"strings"
)

type SearchParam struct {
	Q          string
	Page       int
	StartPrice int
	EndPrice   int
	Sort       string
}

func NewSearchParam(q, sort string, page, startPrice, endPrice int) SearchParam {
	return SearchParam{
		Q:          q,
		Page:       page,
		StartPrice: startPrice,
		EndPrice:   endPrice,
		Sort:       sort,
	}
}

func (p SearchParam) ToQuery() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("q=%s", p.Q))

	if p.Page != 0 {
		b.WriteString("&")
		b.WriteString(fmt.Sprintf("page=%d", p.Page))
	}

	if p.StartPrice != 0 {
		b.WriteString("&")
		b.WriteString(fmt.Sprintf("start_price=%d", p.StartPrice))
	}

	if p.EndPrice != 0 {
		b.WriteString("&")
		b.WriteString(fmt.Sprintf("end_price=%d", p.EndPrice))
	}

	if p.Sort != "" {
		b.WriteString("&")
		b.WriteString(fmt.Sprintf("sort=%s", p.Sort))
	}

	return b.String()
}
