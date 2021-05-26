package models

import (
	"fmt"
)

type Feedback struct {
	Id          uint64 `json:"id"`
	UserId      uint64 `json:"user_id,"`
	ClassroomId uint64 `json:"classroom_id"`
	Comment     string `json:"comment,omitempty"`
}

func (f *Feedback) ObjectId() uint64 {
	return f.Id
}

func (f *Feedback) String() string {
	return fmt.Sprintf("id: %v, user: %v, classroom: %v, comment: %v",
		f.Id, f.UserId, f.ClassroomId, f.Comment)
}
