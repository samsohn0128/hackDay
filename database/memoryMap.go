package database

import (
	"errors"
	"toyHome/entities"
)

var ErrNokey = errors.New("No Key")

var Appliances = map[string]entities.DiscoveredAppliance{
	"device-001": {
		ApplianceId:         "device-001",
		ManufacturerName:    "device-manufacturer-name",
		ModelName:           "smart light",
		Version:             "v1.0",
		FriendlyName:        "livingroom light",
		FriendlyDescription: "controllable light",
		IsIr:                false,
		IsReachable:         true,
		Actions: []string{
			"DecrementBrightness",
			"HealthCheck",
			"IncrementBrightness",
			"SetBrightness",
			"TurnOn",
			"TurnOff",
		},
		ApplianceTypes: []string{
			"LIGHT",
		},
		AdditionalApplianceDetails: map[string]interface{}{},
		Location:                   "",
	},
	"device-002": {
		ApplianceId:         "device-002",
		ManufacturerName:    "device-manufacturer-name",
		ModelName:           "samrt boiler",
		Version:             "v1.0",
		FriendlyName:        "livingroom bolier",
		FriendlyDescription: "controlable boiler",
		IsIr:                false,
		IsReachable:         true,
		Actions: []string{
			"TurnOn",
			"TurnOff",
		},
		ApplianceTypes: []string{
			"Boiler",
		},
		AdditionalApplianceDetails: map[string]interface{}{},
		Location:                   "",
	},
	"device-003": {
		ApplianceId:         "device-003",
		ManufacturerName:    "device-manufacturer-name",
		ModelName:           "smart refri",
		Version:             "v1.0",
		FriendlyName:        "refri",
		FriendlyDescription: "controlable refri",
		IsIr:                false,
		IsReachable:         true,
		Actions: []string{
			"TurnOn",
			"TurnOff",
		},
		ApplianceTypes: []string{
			"Refri",
		},
		AdditionalApplianceDetails: map[string]interface{}{},
		Location:                   "",
	},
	"device-004": {
		ApplianceId:         "device-004",
		ManufacturerName:    "device-manufacturer-name",
		ModelName:           "smart switch",
		Version:             "v1.0",
		FriendlyName:        "switch",
		FriendlyDescription: "controlable switch",
		IsIr:                false,
		IsReachable:         true,
		Actions: []string{
			"TurnOn",
			"TurnOff",
		},
		ApplianceTypes: []string{
			"Switch",
		},
		AdditionalApplianceDetails: map[string]interface{}{},
		Location:                   "",
	},
	"device-005": {
		ApplianceId:         "device-005",
		ManufacturerName:    "device-manufacturer-name",
		ModelName:           "smart aircon",
		Version:             "v1.0",
		FriendlyName:        "livingroom aircon",
		FriendlyDescription: "controlable aircon",
		IsIr:                false,
		IsReachable:         true,
		Actions: []string{
			"TurnOn",
			"TurnOff",
		},
		ApplianceTypes: []string{
			"Aircon",
		},
		AdditionalApplianceDetails: map[string]interface{}{},
		Location:                   "",
	},
}

type memoryMap struct {
}

func InitForTest() {
	instance = &memoryMap{}
}

func (s *memoryMap) GetDiscoveredAppliances(accessToken string) ([]entities.DiscoveredAppliance, error) {
	var dis []entities.DiscoveredAppliance
	for _, value := range Appliances {
		dis = append(dis, value)
	}
	return dis, nil
}

func (s *memoryMap) GetAppliance(accessToken, applianceId string) (entities.DiscoveredAppliance, error) {
	dis, ok := Appliances[applianceId]
	if !ok {
		return dis, ErrNokey
	}
	return dis, nil
}
