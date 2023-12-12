package utils

import (
	"fmt"
	"net/http"
)

const (
	BaseUrl string = "https://api.api-ninjas.com/v1/"
)

type Client struct {
	ApiKey string
}

func Init(apiKey string) Client {
	return Client{
		ApiKey: apiKey,
	}
}

func MakeRequest(url string, apiKey string) (http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return http.Response{}, err
	}

	req.Header.Set("X-Api-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return http.Response{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return http.Response{}, fmt.Errorf("api request failed with status: %v", resp.StatusCode)
	}

	return *resp, nil
}
