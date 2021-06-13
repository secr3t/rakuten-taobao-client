package model

type Detail struct {
	Result struct {
		Item struct {
			Title    string               `json:"title"`
			Images   []string             `json:"images"`
			DescImgs []string             `json:"desc_imgs"`
			NumIid   string               `json:"num_iid"`
			Skus     map[string]SkuDetail `json:"skus"`
			SkuBase  struct {
				Skus []struct {
					PropPath string `json:"propPath"`
					SkuID    string `json:"skuId"`
				} `json:"skus"`
				Prop []struct {
					Values []struct {
						Vid  string `json:"vid"`
						Name string `json:"name"`
					} `json:"values"`
					Name string `json:"name"`
					Pid  string `json:"pid"`
				} `json:"prop"`
			} `json:"sku_base"`
			DescUrl   string `json:"desc_url"`
			DetailUrl string `json:"detail_url"`
		} `json:"item"`
		Status DetailStatus `json:"status"`
	} `json:"result"`
}

type SkuDetail struct {
	PromotionPrice string `json:"promotion_price"`
	Quantity       string `json:"quantity"`
	Price          string `json:"price"`
}

type DetailStatus struct {
	Msg           string `json:"msg"`
	ExecutionTime string `json:"execution_time"`
	Code          int    `json:"code"`
}

func (d Detail) IsSuccess() bool {
	return d.Result.Status.Msg == "success"
}
