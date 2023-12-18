package iban

import (
	"encoding/json"
	"fmt"

	"github.com/nemzyxt/apininjas/utils"
)

const (
	endpoint = utils.BaseUrl + "iban?iban="
)

type IBANClient struct {
	utils.Client
}

type Response struct {
	IBAN          string
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	Country       string
	Checksum      string
	Valid         bool
	BBAN          string
}

func NewClient(apiKey string) IBANClient {
	return IBANClient{
		Client: utils.Init(apiKey),
	}
}

func (c *IBANClient) LookupIBAN(iban string) (Response, error) {
	url := endpoint + fmt.Sprint(iban)

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
