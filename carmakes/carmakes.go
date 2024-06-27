package carmakes

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "carsmakes"
)

type CarMakesClient struct {
	utils.Client
}

type CarMakes struct {
	Makes []string
}

func NewClient(apiKey string) CarMakesClient {
	return CarMakesClient{
		Client: utils.Init(apiKey),
	}
}

func (c *CarMakesClient) GetCarMakes() (CarMakes, error) {
	resp, err := utils.MakeRequest("GET", endpoint, c.ApiKey, nil)
	if err != nil {
		return CarMakes{}, err
	}

	var carmakes CarMakes
	if err := json.NewDecoder(resp.Body).Decode(&carmakes); err != nil {
		return CarMakes{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return carmakes, nil
}
