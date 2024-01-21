package commodityprice

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "commodityprice?name="
)

var (
	commodities = []string{
		"gold futures",
		"soybean oil futures",
		"wheat futures",
		"platinum",
		"micro silver futures",
		"lean hogs futures",
		"corn futures",
		"oat futures",
		"aluminum futures",
		"soybean meal futures",
		"silver futures",
		"lumber futures",
		"live cattle futures",
		"cattle",
		"sugar",
		"natural gas",
		"crude oil",
		"orange juice",
		"coffee",
		"cotton",
		"copper",
		"micro gold futures",
		"feeder cattle futures",
		"rough rice futures",
		"palladium",
		"cocoa",
		"brent crude oil",
		"gasoline rbob",
		"heating oil",
		"class iii milk futures",
	}
)

type CommodityPriceClient struct {
	utils.Client
}

type CommodityPrice struct {
	Exchange string
	Name     string
	Price    float32
	Updated  int32
}

func NewClient(apiKey string) CommodityPriceClient {
	return CommodityPriceClient{
		Client: utils.Init(apiKey),
	}
}

func isSupported(commodity string) bool {
	for _, element := range commodities {
		if strings.ToLower(commodity) == element {
			return true
		}
	}
	return false
}

func (c *CommodityPriceClient) GetCommodityPrice(commodity string) (CommodityPrice, error) {
	if !isSupported(commodity) {
		return CommodityPrice{}, fmt.Errorf("unsupported commodity: %s", commodity)
	}

	url := endpoint + fmt.Sprint(commodity)

	resp, err := utils.MakeRequest("GET", url, c.ApiKey, nil)
	if err != nil {
		return CommodityPrice{}, err
	}

	var commodityprice CommodityPrice
	if err := json.NewDecoder(resp.Body).Decode(&commodityprice); err != nil {
		return CommodityPrice{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return commodityprice, nil
}
