package jokes

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "jokes?limit="
)

type JokesClient struct {
	utils.Client
}

type Joke struct {
	Joke string
}

func NewClient(apiKey string) JokesClient {
	return JokesClient{
		Client: utils.Init(apiKey),
	}
}

func (c *JokesClient) GetJoke() (Joke, error) {
	jokes, err := c.GetJokes(1)
	if err != nil {
		return Joke{}, err
	}
	return jokes[0], nil
}

func (c *JokesClient) GetJokes(limit int) ([]Joke, error) {
	url := endpoint + fmt.Sprint(limit)

	resp, err := utils.MakeRequest(url, c.ApiKey)
	if err != nil {
		return nil, err
	}

	var jokes []Joke
	if err := json.NewDecoder(resp.Body).Decode(&jokes); err != nil {
		return nil, fmt.Errorf("error decoding api response: %v", err)
	}

	return jokes, nil
}
