package cars

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "cars"
)

type CarsClient struct {
	utils.Client
}

type Response struct {
	CityFuelConsumption     float32 `json:"city_mpg"`
	Class                   string  `json:"class"`
	CombinedFuelConsumption float32 `json:"combination_mpg"`
	Cylinders               uint32  `json:"cylinders"`
	Displacement            float32 `json:"displacement"`
	Drive                   string  `json:"drive"`
	FuelType                string  `json:"fuel_type"`
	HighwayFuelConsumption  float32 `json:"highway_mpg"`
	Make                    string  `json:"make"`
	Model                   string  `json:"model"`
	Transmission            string  `json:"transmission"`
	Year                    uint32  `json:"year"`
}

type CarsProps struct {
	Make                       string `json:"make"`
	Model                      string `json:"model"`
	FuelType                   string `json:"fuel_type"`
	Drive                      string `json:"drive"`
	Cylinders                  string `json:"cylinders"`
	Transmission               string `json:"transmission"`
	Year                       string `json:"year"`
	MinCityFuelConsumption     string `json:"min_city_mpg"`
	MaxCityFuelConsumption     string `json:"max_city_mpg"`
	MinHighwayFuelConsumption  string `json:"min_hwy_mpg"`
	MaxHighwayFuelConsumption  string `json:"max_hwy_mpg"`
	MinCombinedFuelConsumption string `json:"min_comb_mpg"` // city and highway
	MaxCombinedFuelConsumption string `json:"max_comb_mpg"` // city and highway
	Limit                      uint32 `json:"limit" default:"5"`
}

func NewClient(apiKey string) CarsClient {
	return CarsClient{
		Client: utils.Init(apiKey),
	}
}

func (c *CarsClient) GetCars(query CarsProps) ([]Response, error) {
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
