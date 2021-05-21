package utils

import "errors"

// MirrorMap creates a new map in a way that values of the old map become
// keys of the new map, and keys of the old map become values of the new map
// If passed map is nil, error will be returned.
// If original map contains two equal values, panic will be invoked
func MirrorMap(dict map[interface{}]interface{}) (map[interface{}]interface{}, error) {
	if dict == nil {
		return nil, errors.New("cannot mirror nil map")
	}

	mirrored := make(map[interface{}]interface{}, len(dict))

	for key, value := range dict {
		if _, present := mirrored[value]; present {
			panic("key collision")
		}
		mirrored[value] = key
	}
	return mirrored, nil
}
