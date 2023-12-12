package bucketlist

import (
	"encoding/json"
	"fmt"
	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "bucketlist"
)

type BucketListClient struct {
	utils.Client
}

type Response struct {
	Item string
}

func NewClient(apiKey string) BucketListClient {
	return BucketListClient{
		Client: utils.Init(apiKey),
	}
}

func (c *BucketListClient) GetBucketListIdea() (Response, error) {
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
