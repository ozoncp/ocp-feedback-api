package models

import (
	"fmt"
)

type Proposal struct {
	Id         uint64 `json:"id"`
	UserId     uint64 `json:"user_id"`
	LessonId   uint64 `json:"lesson_id"`
	DocumentId uint64 `json:"document_id"`
}

func (f *Proposal) ObjectId() uint64 {
	return f.Id
}

func (p *Proposal) String() string {
	return fmt.Sprintf("id: %v, user: %v, lesson: %v, document: %v", p.Id, p.UserId, p.LessonId, p.DocumentId)
}
