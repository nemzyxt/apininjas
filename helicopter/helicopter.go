package helicopter

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "helicopter"
)

type HelicopterClient struct {
	utils.Client
}

type Response struct {
	Manufacturer      string `json:"manufacturer"`
	Model             string `json:"model"`
	MaxSpeed          string `json:"max_speed_sl_knots"`
	CruiseSpeed       string `json:"cruise_speed_sl_knots"`
	Range             string `json:"range_nautical_miles"`
	CruiseTime        string `json:"cruise_time_min"`
	FuelCapacity      string `json:"fuel_capacity_gallons"`
	FuelConsumption   string `json:"fuel_consumption_gallons_pr_hr"`
	GrossExternalLoad string `json:"gross_external_load_lbs"`
	ExternalLoadLimit string `json:"external_load_limit_lbs"`
	MainRotorDiameter string `json:"main_rotor_diameter_ft"`
	NumBlades         string `json:"num_blades"`
	BladeMaterial     string `json:"blade_material"`
	RotorType         string `json:"rotor_type"`
	StorageWidth      string `json:"storage_width_ft"`
	Length            string `json:"length_ft"`
	Height            string `json:"height_ft"`
}

type HelicopterProps struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	MinSpeed     string `json:"min_speed"`
	MaxSpeed     string `json:"max_speed"`
	MinRange     string `json:"min_range"`
	MaxRange     string `json:"max_range"`
	MinLength    string `json:"min_length"`
	MaxLength    string `json:"max_length"`
	MinHeight    string `json:"min_height"`
	MaxHeight    string `json:"max_height"`
	Limit        uint32 `json:"limit" default:"1"`
}

func NewClient(apiKey string) HelicopterClient {
	return HelicopterClient{
		Client: utils.Init(apiKey),
	}
}

func (c *HelicopterClient) GetHelicopters(query HelicopterProps) ([]Response, error) {
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
