package vinlookup

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "vinlookup?vin="
)

type VINNumberClient struct {
	utils.Client
}

type Response struct {
	VIN          string
	Country      string
	Manufacturer string
	Region       string
	WMI          string
	VDS          string
	VIS          string
	Years        []int
}

func NewClient(apiKey string) VINNumberClient {
	return VINNumberClient{
		Client: utils.Init(apiKey),
	}
}

func (c *VINNumberClient) LookupVIN(vin string) (Response, error) {
	if len(vin) != 17 {
		return Response{}, fmt.Errorf("provide a valid 17-character vin")
	}

	url := endpoint + fmt.Sprint(vin)

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
