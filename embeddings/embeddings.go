package embeddings

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "embeddings"
)

type EmbeddingsClient struct {
	utils.Client
}

type RequestBody struct {
	Text string `json:"text"`
}

type Response struct {
	Embeddings []float64
}

func NewClient(apiKey string) EmbeddingsClient {
	return EmbeddingsClient{
		Client: utils.Init(apiKey),
	}
}

func (c *EmbeddingsClient) EncodeText(text string) (Response, error) {
	request_body := RequestBody{
		Text: text,
	}

	body, err := json.Marshal(request_body)
	if err != nil {
		return Response{}, err
	}

	resp, err := utils.MakeRequest("POST", endpoint, c.ApiKey, body)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return response, nil
}
