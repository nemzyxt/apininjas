package airlines

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "airlines"
)

type AirlinesClient struct {
	utils.Client
}

type Response struct {
	IATA    string            `json:"iata"`
	ICAO    string            `json:"icao"`
	Fleet   map[string]uint32 `json:"fleet"`
	LogoURL string            `json:"logo_url"`
	Name    string            `json:"name"`
}

type AirlineProps struct {
	ICAO string `json:"icao"`
	IATA string `json:"iata"`
	Name string `json:"name"`
}

func NewClient(apiKey string) AirlinesClient {
	return AirlinesClient{
		Client: utils.Init(apiKey),
	}
}

func (c *AirlinesClient) GetAirlines(query AirlineProps) ([]Response, error) {
	queryString, err := utils.MakeQueryString(query)
	if err != nil {
		return []Response{}, err
	}
	url := endpoint + "?" + queryString

	resp, err := utils.MakeRequest("GET", url, c.ApiKey, nil)
	if err != nil {
		return []Response{}, err
	}

	var response []Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return []Response{}, fmt.Errorf("error decoding api response: %v", err)
	}
	resp.Body.Close()
	return response, nil
}
