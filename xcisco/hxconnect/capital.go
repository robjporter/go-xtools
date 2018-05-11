package hxconnect

import (
	"../request"
)

type capurls struct {
	LicenseAgreement string
	ValidateContract string
	SubmitData       string
}

type capital struct {
	Company       string
	Start         string
	Duration      string
	EncryptionKey string
	Nodes         int
	Costs         interface{}
}

func getCapURLS() *capurls {
	return &capurls{
		LicenseAgreement: "/licenseagreement",
		ValidateContract: "/validatecontract",
		SubmitData:       "/submitdata",
	}
}

func (c *Connection) GetLicenseAgreement() (*request.Client,error) {
	req,e := c.sendGetRequest(c.Metrics.Server, c.capurls.LicenseAgreement)
	return req,e
}

func (c *Connection) ValidateContactNumberWithCisco(contract string) (*request.Client,error)  {
	data := make(map[string]string)
	data["contractnumber"] = contract
	req,e := c.sendPostRequest(c.Metrics.Server, c.capurls.ValidateContract, data)
	return req,e
}

func (c *Connection) SendDataToCaptial(json string) (*request.Client,error) {
	req, e := c.sendPostRequestSimple(c.Metrics.Server, c.capurls.SubmitData, json)
	return req,e
}
