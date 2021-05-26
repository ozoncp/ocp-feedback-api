package utils

import "github.com/ozoncp/ocp-feedback-api/internal/models"

type void struct{}

// FilterAllowed filters out items from list using blockList as a filter.
// That is, only the items which don't belong to the blockList will be returned.
// If one of arguments is nil, panic will be invoked
func FilterAllowed(list []models.Entity, blockList []uint64) ([]models.Entity, error) {
	if list == nil || blockList == nil {
		panic("cannot filter nil slice")
	}

	allowed := []models.Entity{}
	blockListSet := make(map[uint64]void, len(blockList))

	for i := 0; i < len(blockList); i++ {
		blockListSet[blockList[i]] = void{}
	}

	for i := 0; i < len(list); i++ {
		if _, exists := blockListSet[list[i].ObjectId()]; !exists {
			allowed = append(allowed, list[i])
		}
	}
	return allowed, nil
}
