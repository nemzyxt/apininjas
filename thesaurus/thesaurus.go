package thesaurus

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "thesaurus?word="
)

type ThesaurusClient struct {
	utils.Client
}

type Response struct {
	Word     string
	Synonyms []string
	Antonyms []string
}

func NewClient(apiKey string) ThesaurusClient {
	return ThesaurusClient{
		Client: utils.Init(apiKey),
	}
}

func (c *ThesaurusClient) LookupWord(word string) (Response, error) {
	url := endpoint + fmt.Sprint(word)

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
