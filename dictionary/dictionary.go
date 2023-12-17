package dictionary

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "dictionary?word="
)

type DictClient struct {
	utils.Client
}

type Response struct {
	Definition string
	Word       string
	Valid      bool
}

func NewClient(apiKey string) DictClient {
	return DictClient{
		Client: utils.Init(apiKey),
	}
}

func (c *DictClient) CheckWord(word string) (Response, error) {
	url := endpoint + word

	resp, err := utils.MakeRequest("GET", url, c.ApiKey, nil)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return response, nil
}
