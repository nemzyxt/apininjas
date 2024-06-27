package carmodels

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "carmodels"
)

type CarModelsClient struct {
	utils.Client
}

type CarModels struct {
	Models []string
}

func NewClient(apiKey string) CarModelsClient {
	return CarModelsClient{
		Client: utils.Init(apiKey),
	}
}

func (c *CarModelsClient) GetCarModels() (CarModels, error) {
	resp, err := utils.MakeRequest("GET", endpoint, c.ApiKey, nil)
	if err != nil {
		return CarModels{}, err
	}

	var carmodels CarModels
	if err := json.NewDecoder(resp.Body).Decode(&carmodels); err != nil {
		return CarModels{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return carmodels, nil
}
