package motorcycles

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "motorcycles"
)

type MotorcyclesClient struct {
	utils.Client
}

type Response struct {
	Make              string `json:"make"`
	Model             string `json:"model"`
	Year              string `json:"year"`
	Type              string `json:"type"`
	Displacement      string `json:"displacement"`
	Engine            string `json:"engine"`
	Power             string `json:"power"`
	Torque            string `json:"torque"`
	Compression       string `json:"compression"`
	BoreStroke        string `json:"bore_stroke"`
	ValvesPerCylinder string `json:"valves_per_cylinder"`
	FuelSystem        string `json:"fuel_system"`
	FuelControl       string `json:"fuel_control"`
	Ignition          string `json:"ignition"`
	Lubrication       string `json:"lubrication"`
	Cooling           string `json:"cooling"`
	Gearbox           string `json:"gearbox"`
	Transmission      string `json:"transmission"`
	Clutch            string `json:"clutch"`
	Frame             string `json:"frame"`
	FrontSuspension   string `json:"front_suspension"`
	FrontWheelTravel  string `json:"front_wheel_travel"`
	RearSuspension    string `json:"rear_suspension"`
	RearWheelTravel   string `json:"rear_wheel_travel"`
	FrontTire         string `json:"front_tire"`
	RearTire          string `json:"rear_tire"`
	FrontBrakes       string `json:"front_brakes"`
	RearBrakes        string `json:"rear_brakes"`
	TotalWeight       string `json:"total_weight"`
	SeatHeight        string `json:"seat_height"`
	TotalHeight       string `json:"total_height"`
	TotalLength       string `json:"total_length"`
	TotalWidth        string `json:"total_width"`
	GroundClearance   string `json:"ground_clearance"`
	Wheelbase         string `json:"wheelbase"`
	FuelCapacity      string `json:"fuel_capacity"`
	Starter           string `json:"starter"`
}

type MotorcycleProps struct {
	Make   string `json:"make"`
	Model  string `json:"model"`
	Year   string `json:"year"`
	Offset string `json:"offset" default:"0"`
}

func NewClient(apiKey string) MotorcyclesClient {
	return MotorcyclesClient{
		Client: utils.Init(apiKey),
	}
}

func (c *MotorcyclesClient) GetMotorcycles(query MotorcycleProps) ([]Response, error) {
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
