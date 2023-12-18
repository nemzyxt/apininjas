package urllookup

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "urllookup?url="
)

type UrlLookupClient struct {
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
	URL         string
}

func NewClient(apiKey string) UrlLookupClient {
	return UrlLookupClient{
		Client: utils.Init(apiKey),
	}
}

func (c *UrlLookupClient) LookupURL(url string) (Response, error) {
	final_url := endpoint + fmt.Sprint(url)

	resp, err := utils.MakeRequest("GET", final_url, c.ApiKey, nil)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("error decoding api response: %v", err)
	}

	return response, nil
}
