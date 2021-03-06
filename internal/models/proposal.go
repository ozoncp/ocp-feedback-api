package models

import (
	"fmt"
	"unsafe"
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

func (p *Proposal) Size() uint64 {
	return uint64(unsafe.Sizeof(p.Id) +
		unsafe.Sizeof(p.UserId) +
		unsafe.Sizeof(p.LessonId) +
		unsafe.Sizeof(p.DocumentId))
}
