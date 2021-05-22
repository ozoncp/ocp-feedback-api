package utils

import (
	"errors"
	"math"

	"github.com/ozoncp/ocp-feedback-api/internal/models/entity"
)

// SplitSlice splits passed slice into equal chunks.
// If the size the slice is not divisible by the chunk size, the last chunk
// will be less than the chunk size
// If slice size is smaller than the chunk size or if the chunk size is zero,
// the whole slice will be returned.
// If passed slice is nil, error will be returned.
// If passed chunk size is negative, error will be returned
func SplitSlice(slice []entity.Entity, chunkSize int) ([][]entity.Entity, error) {
	if slice == nil {
		return nil, errors.New("cannot split nil slice")
	}
	if chunkSize < 0 {
		return nil, errors.New("cannot split slice into chunks of negative size")
	}
	if chunkSize == 0 {
		return [][]entity.Entity{slice}, nil
	}

	chunks := [][]entity.Entity{}
	right_bound := 0

	for i := 0; i < len(slice); i += chunkSize {

		// integer overflow is possible if chunk size and len(slice) are huge
		if math.MaxInt32-right_bound < chunkSize || right_bound+chunkSize > len(slice) {
			right_bound = len(slice)
		} else {
			right_bound += chunkSize
		}

		chunks = append(chunks, slice[i:right_bound])
	}
	return chunks, nil
}
