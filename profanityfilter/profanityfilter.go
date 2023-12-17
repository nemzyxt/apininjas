package profanityfilter

import (
	"encoding/json"
	"fmt"
	"github.com/nemzyxt/apininjas/utils"
	"net/url"
)

const (
	endpoint = utils.BaseUrl + "profanityfilter?text="
)

type ProfanityFilterClient struct {
	utils.Client
}

type Response struct {
	Original     string
	Censored     string
	HasProfanity bool `json:"has_profanity"`
}

func NewClient(apiKey string) ProfanityFilterClient {
	return ProfanityFilterClient{
		Client: utils.Init(apiKey),
	}
}

func (c *ProfanityFilterClient) RunFilter(text string) (Response, error) {
	url := endpoint + url.QueryEscape(text)

	resp, err := utils.MakeRequest(url, c.ApiKey)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return response, nil
}
