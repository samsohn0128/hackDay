package database

import "toyHome/entities"

var instance Querier

type Querier interface {
	GetDiscoveredAppliances(accessToken string) ([]entities.DiscoveredAppliance, error)
	GetAppliance(accessToken, applianceId string) (entities.DiscoveredAppliance, error)
}

func GetInstance() Querier {
	if instance == nil {
		return &memoryMap{}
	}
	return instance
}
