package utils

import "errors"

type void struct{}

// FilterAllowed filters out items from list using blockList as a filter.
// That is, only the items which don't belong to the blockList will be returned.
// If one of arguments is nil, error will be returned
func FilterAllowed(list []interface{}, blockList []interface{}) ([]interface{}, error) {
	if list == nil || blockList == nil {
		return nil, errors.New("cannot filter nil slice")
	}

	allowed := []interface{}{}
	blockListSet := make(map[interface{}]void, len(blockList))

	for i := 0; i < len(blockList); i++ {
		blockListSet[blockList[i]] = void{}
	}

	for i := 0; i < len(list); i++ {
		if _, exists := blockListSet[list[i]]; !exists {
			allowed = append(allowed, list[i])
		}
	}
	return allowed, nil
}
