package aircraft

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "aircraft"
)

type AircraftClient struct {
	utils.Client
}

type Response struct {
	Manufacturer      string `json:"manufacturer"`
	Model             string `json:"model"`
	EngineType        string `json:"engine_type"`
	EngineThrust      string `json:"engine_thrust_lb_ft"`
	MaxSpeed          string `json:"max_speed_knots"`
	CruiseSpeed       string `json:"cruise_speed_knots"`
	Ceiling           string `json:"ceiling_ft"`
	TakeoffGroundRun  string `json:"takeoff_ground_run_ft"`
	LandingGroundRoll string `json:"landing_ground_roll_ft"`
	GrossWeight       string `json:"gross_weight_lbs"`
	EmptyWeight       string `json:"empty_weight_lbs"`
	Length            string `json:"length_ft"`
	Height            string `json:"height_ft"`
	WingSpan          string `json:"wing_span_ft"`
	Range             string `json:"range_nautical_miles"`
}

type AircraftProps struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	EngineType   string `json:"engine_type" default:"piston"`
	MinSpeed     string `json:"min_speed"`
	MaxSpeed     string `json:"max_speed"`
	MinRange     string `json:"min_range"`
	MaxRange     string `json:"max_range"`
	MinLength    string `json:"min_length"`
	MaxLength    string `json:"max_length"`
	MinHeight    string `json:"min_height"`
	MaxHeight    string `json:"max_height"`
	MinWingspan  string `json:"min_wingspan"`
	Limit        uint32 `json:"limit" default:"1"`
}

func NewClient(apiKey string) AircraftClient {
	return AircraftClient{
		Client: utils.Init(apiKey),
	}
}

func (c *AircraftClient) GetAircraft(query AircraftProps) ([]Response, error) {
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
