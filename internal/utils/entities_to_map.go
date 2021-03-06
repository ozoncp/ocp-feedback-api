package utils

import (
	"errors"

	"github.com/ozoncp/ocp-feedback-api/internal/models"
)

// EntitiesToMap converts slice of entities into a map
// where key is an entity id and value is an entity itself
// If passed slice containts two equal ids, panic will be invoked
func EntitiesToMap(entities []models.Entity) (map[uint64]models.Entity, error) {
	if entities == nil {
		return nil, errors.New("cannot convert nil slice")
	}

	entitiesMap := make(map[uint64]models.Entity, len(entities))

	for i := 0; i < len(entities); i++ {
		e := entities[i]
		if _, present := entitiesMap[e.ObjectId()]; present {
			panic("id collision")
		}
		entitiesMap[e.ObjectId()] = e
	}
	return entitiesMap, nil
}
