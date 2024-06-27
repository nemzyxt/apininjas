package loremipsum

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "loremipsum"
)

type LoremIpsumClient struct {
	utils.Client
}

type Response struct {
	Text string `json:"text"`
}

type Params struct {
	MaxLength           string `json:"max_length"`
	Paragraphs          string `json:"paragraphs"`
	StartWithLoremIpsum bool   `json:"start_with_lorem_ipsum"`
	Random              bool   `json:"random"`
}

func NewClient(apiKey string) LoremIpsumClient {
	return LoremIpsumClient{
		Client: utils.Init(apiKey),
	}
}

func (c *LoremIpsumClient) GetLoremIpsumText(query Params) ([]Response, error) {
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
