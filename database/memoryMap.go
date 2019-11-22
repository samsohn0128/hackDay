package database

import (
	"encoding/json"
	"errors"
	"toyHome/entities"
)

type memoryMap struct {
}

func GetInstances() Querier {
	return &memoryMap{}
}

func (s *memoryMap) GetDiscoveredAppliances(accessToken string) ([]entities.DiscoveredAppliance, error) {
	var dis entities.DiscoveredAppliances
	err := json.Unmarshal([]byte(`{
  "discoveredAppliances": [
    {
      "applianceId": "device-001",
      "manufacturerName": "device-manufacturer-name",
      "modelName": "스마트 전등",
      "version": "v1.0",
      "friendlyName": "거실 전등",
      "friendlyDescription": "스마트폰으로 제어할 수 있는 전등",
      "isIr": false,
      "isReachable": true,
      "actions": [
        "DecrementBrightness",
        "HealthCheck",
        "IncrementBrightness",
        "SetBrightness",
        "TurnOn",
        "TurnOff"
      ],
      "applianceTypes": [
        "LIGHT"
      ],
      "additionalApplianceDetails": {},
      "location": ""
    },
    {
      "applianceId": "device-002",
      "manufacturerName": "device-manufacturer-name",
      "modelName": "스마트 플러그",
      "version": "v1.0",
      "friendlyName": "부엌 플러그",
      "friendlyDescription": "에너지를 절약하는 플러그",
      "isIr": false,
      "isReachable": true,
      "actions": [
        "HealthCheck",
        "TurnOn",
        "TurnOff"
      ],
      "applianceTypes": [
        "SMARTPLUG"
      ],
      "additionalApplianceDetails": {},
      "location": "LIVING_ROOM"
    }
  ]
}`), &dis)
	if err != nil {
		return nil, err
	}
	return dis.DiscoveredAppliances, nil
}

func (s *memoryMap) GetAppliance(accessToken string) ([]entities.DiscoveredAppliance, error) {
	//var dis entities.DiscoveredAppliances
	//err := json.Unmarshal()
	return nil, errors.New("Not Imple.")
}
