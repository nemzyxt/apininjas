package randomuser

import (
	"encoding/json"
	"fmt"
	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "randomuser"
)

type RandomUserClient struct {
	utils.Client
}

type Response struct {
	Username string
	Sex      string
	Address  string
	Name     string
	Email    string
	Birthday string
}

func NewClient(apiKey string) RandomUserClient {
	return RandomUserClient{
		Client: utils.Init(apiKey),
	}
}

func (c *RandomUserClient) GetRandomUser() (Response, error) {
	url := endpoint

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
