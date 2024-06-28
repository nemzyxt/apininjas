package randomword

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "randomword"
)

type RandomWordClient struct {
	utils.Client
}

type Response struct {
	Word []string `json:"word"`
}

type RandomWordProps struct {
	Type  string `json:"type"`
	Limit string `json:"limit" default:"1"`
}

func NewClient(apiKey string) RandomWordClient {
	return RandomWordClient{
		Client: utils.Init(apiKey),
	}
}

func (c *RandomWordClient) GetRandomWord(query RandomWordProps) (Response, error) {
	queryString, err := utils.MakeQueryString(query)
	if err != nil {
		return Response{}, err
	}
	url := endpoint + "?" + queryString

	resp, err := utils.MakeRequest("GET", url, c.ApiKey, nil)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("error decoding api response: %v", err)
	}
	resp.Body.Close()
	return response, nil
}
