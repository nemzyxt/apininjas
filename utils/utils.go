package utils

import (
	"bytes"
	"fmt"
	"io"
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

func MakeRequest(method string, url string, apiKey string, body []byte) (http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return http.Response{}, err
	}

	req.Header.Set("X-Api-Key", apiKey)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
		req.Body = io.NopCloser(bytes.NewBuffer(body))
		req.ContentLength = int64(len(body))
	}

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
