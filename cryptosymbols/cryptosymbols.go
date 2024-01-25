package cryptosymbols

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "cryptosymbols"
)

type CryptoSymbolsClient struct {
	utils.Client
}

type CryptoSymbols struct {
	Symbols []string
}

func NewClient(apiKey string) CryptoSymbolsClient {
	return CryptoSymbolsClient{
		Client: utils.Init(apiKey),
	}
}

func (c *CryptoSymbolsClient) GetCryptoSymbols() (CryptoSymbols, error) {
	resp, err := utils.MakeRequest("GET", endpoint, c.ApiKey, nil)
	if err != nil {
		return CryptoSymbols{}, err
	}

	var cryptosymbols CryptoSymbols
	if err := json.NewDecoder(resp.Body).Decode(&cryptosymbols); err != nil {
		return CryptoSymbols{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return cryptosymbols, nil
}
