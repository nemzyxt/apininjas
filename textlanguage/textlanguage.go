package textlanguage

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "textlanguage?text="
)

type TextLanguageClient struct {
	utils.Client
}

type Response struct {
	ISO      string
	Language string
}

func NewClient(apiKey string) TextLanguageClient {
	return TextLanguageClient{
		Client: utils.Init(apiKey),
	}
}

func (c *TextLanguageClient) IdentifyLanguage(text string) (Response, error) {
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
