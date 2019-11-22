package entities

import "encoding/json"

type Appliance struct {
	ApplianceId string `json:"applianceId"`
}

type RequestPayload struct {
	AccessToken string    `json:"accessToken"`
	Appliances  Appliance `json:"appliance"`
}

type Payload struct {
	AccessToken string `json:"accessToken"`
}

type Request struct {
	Head    Header          `json:"header"`
	Payload json.RawMessage `json:"payload"`
}

type Header struct {
	MessageId      string `json:"messageId"`
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	PayloadVersion string `json:"payloadVersion"`
}

type Response Request

type DiscoveredAppliances struct {
	DiscoveredAppliances []DiscoveredAppliance `json:"discoveredAppliances"`
}

type DiscoveredAppliance struct {
	ApplianceId                string      `json:"applianceId"`
	ManufacturerName           string      `json:"manufacturerName"`
	ModelName                  string      `json:"modelName"`
	Version                    string      `json:"version"`
	FriendlyName               string      `json:"friendlyName"`
	FriendlyDescription        string      `json:"friendlyDescription"`
	IsIr                       bool        `json:"isIr"`
	IsReachable                bool        `json:"isReachable"`
	Actions                    []string    `json:"actions"`
	ApplianceTypes             []string    `json:"applianceTypes"`
	AdditionalApplianceDetails interface{} `json:"additionalApplianceDetails"`
	Location                   string      `json:"location"`
}
