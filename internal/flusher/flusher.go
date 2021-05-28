package flusher

import "github.com/ozoncp/ocp-feedback-api/internal/models"

type Flusher interface {
	Flush(entities []models.Entity) []models.Entity
}
