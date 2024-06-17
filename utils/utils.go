package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func MakeQueryString(query interface{}) (string, error) {
	jsonData, err := json.Marshal(query)
	if err != nil {
		return "", err
	}
	var data map[string]interface{}
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return "", err
	}

	values := url.Values{}

	// Convert the map to query parameters
	for key, value := range data {
		switch v := value.(type) {
		case string:
			values.Set(key, v)
		case int, int8, int16, int32, int64, float32, float64:
			values.Set(key, fmt.Sprintf("%v", v))
		case bool:
			values.Set(key, fmt.Sprintf("%t", v))
		default:
			// For complex types, we might need additional handling
			return "", fmt.Errorf("unsupported type for key %s", key)
		}
	}

	// Encode the values into a query string
	return values.Encode(), nil
}
