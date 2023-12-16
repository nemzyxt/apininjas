package facts

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "facts?limit="
)

type FactsClient struct {
	utils.Client
}

type Fact struct {
	Fact string
}

func NewClient(apiKey string) FactsClient {
	return FactsClient{
		Client: utils.Init(apiKey),
	}
}

func (c *FactsClient) GetFact() (Fact, error) {
	facts, err := c.GetFacts(1)
	if err != nil {
		return Fact{}, err
	}
	return facts[0], nil
}

func (c *FactsClient) GetFacts(limit int) ([]Fact, error) {
	url := endpoint + fmt.Sprint(limit)

	resp, err := utils.MakeRequest(url, c.ApiKey)
	if err != nil {
		return nil, err
	}

	var facts []Fact
	if err := json.NewDecoder(resp.Body).Decode(&facts); err != nil {
		return nil, fmt.Errorf("error decoding api response: %v", err)
	}

	return facts, nil
}
