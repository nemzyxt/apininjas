package dictionary

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	PATH = "dictionary?word="
)

type DictClient struct {
	utils.Client
}

type Response struct {
	Definition string
	Word string
	Valid bool
}

func NewClient(apiKey string) DictClient {
	return DictClient{
		Client: utils.Init(apiKey),
	}
}

func (c *DictClient) CheckWord(word string) (Response, error) {
	url := utils.BaseUrl + PATH + word

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
			return Response{}, err
	}

	req.Header.Set("X-Api-Key", c.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
			return Response{}, fmt.Errorf("api request failed with status: %v", resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
			return Response{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return response, nil
}
