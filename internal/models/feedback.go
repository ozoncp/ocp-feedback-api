package models

import (
	"fmt"
	"unsafe"
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

func (f *Feedback) Size() uint64 {
	sz := unsafe.Sizeof(f.Id) +
		unsafe.Sizeof(f.ClassroomId) +
		unsafe.Sizeof(f.UserId) +
		(unsafe.Sizeof(f.Comment))
	return uint64(len(f.Comment)) + uint64(sz)
}
