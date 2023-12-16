package rhyme

import (
	"encoding/json"
	"fmt"
	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "rhyme?word="
)

type RhymeClient struct {
	utils.Client
}

func NewClient(apiKey string) RhymeClient {
	return RhymeClient{
		Client: utils.Init(apiKey),
	}
}

func (c *RhymeClient) GetRhymingWords(word string) ([]string, error) {
	url := endpoint + fmt.Sprint(word)

	resp, err := utils.MakeRequest(url, c.ApiKey)
	if err != nil {
		return []string{}, err
	}

	var words []string
	if err := json.NewDecoder(resp.Body).Decode(&words); err != nil {
		return []string{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return words, nil
}
