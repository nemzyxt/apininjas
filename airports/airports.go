package airports

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "airports"
)

type AirportsClient struct {
	utils.Client
}

type Response struct {
	ICAO      string `json:"icao"`
	IATA      string `json:"iata"`
	Name      string `json:"name"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	Elevation string `json:"elevation_ft"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Timezone  string `json:"timezone"`
}

type AirportProps struct {
	ICAO         string `json:"icao"`
	IATA         string `json:"iata"`
	Name         string `json:"name"`
	Country      string `json:"country"`
	Region       string `json:"region"`
	City         string `json:"city"`
	Timezone     string `json:"timezone"`
	MinElevation string `json:"min_elevation"`
	MaxElevation string `json:"max_elevation"`
	Offset       string `json:"offset"`
}

func NewClient(apiKey string) AirportsClient {
	return AirportsClient{
		Client: utils.Init(apiKey),
	}
}

func (c *AirportsClient) GetAirports(query AirportProps) ([]Response, error) {
	queryString, err := utils.MakeQueryString(query)
	if err != nil {
		return []Response{}, err
	}
	url := endpoint + "?" + queryString

	resp, err := utils.MakeRequest("GET", url, c.ApiKey, nil)
	if err != nil {
		return []Response{}, err
	}

	var response []Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return []Response{}, fmt.Errorf("error decoding api response: %v", err)
	}
	resp.Body.Close()
	return response, nil
}
