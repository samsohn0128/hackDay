package api

import (
	"encoding/json"
	"errors"
	"toyHome/database"
	"toyHome/entities"
)

func ResponseDiscovery(message json.RawMessage) ([]byte, error) {
	var payLoad entities.Payload
	err := json.Unmarshal(message, &payLoad)
	if err != nil{
		return nil,err
	}

	discoveredAppliance, err := database.GetInstances().GetDiscoveredAppliances(payLoad.AccessToken)
	if err != nil{
		return nil, err
	}
	var discoveredAppliances entities.DiscoveredAppliances
	discoveredAppliances.DiscoveredAppliances = discoveredAppliance
	out, err := json.Marshal(discoveredAppliances)
	if err != nil{
		return nil,err
	}
	return out,nil
}

func ResponseRequest(message json.RawMessage) ([]byte, error) {
	var requestPayload entities.RequestPayload
	err := json.Unmarshal(message, &requestPayload)
	if err != nil{
		return nil, err
	}
	//requestPayload, err := database.GetInstances().
	return nil, errors.New("Not imple.")
}
