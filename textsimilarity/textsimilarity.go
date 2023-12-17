package textsimilarity

import (
	"encoding/json"
	"fmt"
	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "textsimilarity"
)

type TextSimilarityClient struct {
	utils.Client
}

type RequestBody struct {
	Text1 string `json:"text_1"`
	Text2 string `json:"text_2"`
}

type Response struct {
	Similarity float64
}

func NewClient(apiKey string) TextSimilarityClient {
	return TextSimilarityClient{
		Client: utils.Init(apiKey),
	}
}

func (c *TextSimilarityClient) ComputeSimilarity(text1, text2 string) (Response, error) {
	url := endpoint
	request_body := RequestBody{
		Text1: text1,
		Text2: text2,
	}

	body, err := json.Marshal(request_body)
	if err != nil {
		return Response{}, err
	}

	resp, err := utils.MakeRequest(url, c.ApiKey, body)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return response, nil
}
