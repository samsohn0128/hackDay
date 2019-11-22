package database

import "toyHome/entities"

type Querier interface {
	GetDiscoveredAppliances(accessToken string) ([]entities.DiscoveredAppliance, error)
}
