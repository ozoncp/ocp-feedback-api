package utils

import (
	"errors"

	"github.com/ozoncp/ocp-feedback-api/internal/models/entity"
)

// EntitiesToMap converts slice of entities into a map
// where key is an entity id and value is an entity itself
// If passed slice is nil, error will be returned
// If passed slice containts two equal ids, panic will be invoked
func EntitiesToMap(entities []entity.Entity) (map[uint64]entity.Entity, error) {
	if entities == nil {
		return nil, errors.New("cannot convert nil slice")
	}

	entitiesMap := make(map[uint64]entity.Entity, len(entities))

	for i := 0; i < len(entities); i++ {
		e := entities[i]
		if _, present := entitiesMap[e.Id()]; present {
			panic("id collision")
		}
		entitiesMap[e.Id()] = e
	}
	return entitiesMap, nil
}
