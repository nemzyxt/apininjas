package iplookup

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "iplookup?address="
)

type IPLookupClient struct {
	utils.Client
}

type Response struct {
	IsValid     bool `json:"is_valid"`
	Country     string
	CountryCode string `json:"country_code"`
	RegionCode  string `json:"region_code"`
	Region      string
	City        string
	Zip         string
	Lat         float64
	Lon         float64
	Timezone    string
	ISP         string
	Address     string
}

func NewClient(apiKey string) IPLookupClient {
	return IPLookupClient{
		Client: utils.Init(apiKey),
	}
}

func (c *IPLookupClient) GetInfo(ip string) (Response, error) {
	url := endpoint + fmt.Sprint(ip)

	resp, err := utils.MakeRequest("GET", url, c.ApiKey, nil)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return response, nil
}
