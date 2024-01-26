package cryptoprice

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "cryptoprice?symbol="
)

type CryptoPriceClient struct {
	utils.Client
}

type CryptoPrice struct {
	Symbol    string
	Price     string
	Timestamp int32
}

func NewClient(apiKey string) CryptoPriceClient {
	return CryptoPriceClient{
		Client: utils.Init(apiKey),
	}
}

func (c *CryptoPriceClient) GetCryptoPrice(symbol string) (CryptoPrice, error) {
	url := endpoint + fmt.Sprint(symbol)

	resp, err := utils.MakeRequest("GET", url, c.ApiKey, nil)
	if err != nil {
		return CryptoPrice{}, err
	}

	var cryptoprice CryptoPrice
	if err := json.NewDecoder(resp.Body).Decode(&cryptoprice); err != nil {
		return CryptoPrice{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return cryptoprice, nil
}
