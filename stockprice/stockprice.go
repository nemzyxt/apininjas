package stockprice

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "stockprice?ticker="
)

type StockPriceClient struct {
	utils.Client
}

type StockPrice struct {
	Ticker   string
	Name     string
	Price    float32
	Exchange string
	Updated  int32
}

func NewClient(apiKey string) StockPriceClient {
	return StockPriceClient{
		Client: utils.Init(apiKey),
	}
}

func (c *StockPriceClient) GetStockPrice(ticker string) (StockPrice, error) {
	url := endpoint + fmt.Sprint(ticker)

	resp, err := utils.MakeRequest("GET", url, c.ApiKey, nil)
	if err != nil {
		return StockPrice{}, err
	}

	var stockprice StockPrice
	if err := json.NewDecoder(resp.Body).Decode(&stockprice); err != nil {
		return StockPrice{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return stockprice, nil
}
