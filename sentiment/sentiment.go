package sentiment

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "sentiment?text="
)

type SentimentClient struct {
	utils.Client
}

type Response struct {
	Score     float64
	Text      string
	Sentiment string
}

func NewClient(apiKey string) SentimentClient {
	return SentimentClient{
		Client: utils.Init(apiKey),
	}
}

func (c *SentimentClient) AnalyzeSentiment(text string) (Response, error) {
	url := endpoint + url.QueryEscape(text)

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
