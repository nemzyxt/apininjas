package riddles

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "riddles?limit="
)

type Riddle struct {
	Title    string
	Question string
	Answer   string
}

type RiddlesClient struct {
	utils.Client
}

func NewClient(apiKey string) RiddlesClient {
	return RiddlesClient{
		Client: utils.Init(apiKey),
	}
}

func (c *RiddlesClient) GetRiddle() (Riddle, error) {
	riddles, err := c.GetRiddles(1)
	if err != nil {
		return Riddle{}, err
	}
	return riddles[0], nil
}

func (c *RiddlesClient) GetRiddles(limit int) ([]Riddle, error) {
	url := endpoint + fmt.Sprint(limit)

	resp, err := utils.MakeRequest("GET", url, c.ApiKey, nil)
	if err != nil {
		return []Riddle{}, err
	}

	var riddles []Riddle
	if err := json.NewDecoder(resp.Body).Decode(&riddles); err != nil {
		return []Riddle{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return riddles, nil
}
