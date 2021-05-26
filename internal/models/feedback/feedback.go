package feedback

import (
	"fmt"

	"github.com/ozoncp/ocp-feedback-api/internal/models/entity"
)

const (
	ErrInvalidId          = FeedbackErr("provided id is invalid")
	ErrInvalidUserId      = FeedbackErr("provided user id is invalid")
	ErrInvalidClassroomId = FeedbackErr("provided classroom id is invalid")
	ErrInvalidComment     = FeedbackErr("provided comment is invalid")

	formatString = "id: %v, user: %v, classroom: %v, comment: %v"
)

// FeedbackErr are errors that happen while interacting with a feedback
type FeedbackErr string

func (e FeedbackErr) Error() string {
	return string(e)
}

type Feedback interface {
	entity.Entity
	ClassroomId() uint64
	UpdateComment(string) error
	Comment() string
}

type feedback struct {
	id          uint64
	userId      uint64
	classroomId uint64
	comment     string
}

// New returns a Feedback object representing a user feedback
// If passed id is invalid, ErrInvalidId error will be returned
// If passed userId is invalid, ErrInvalidUserId error will be returned
// If passed classroomId is invalid, ErrInvalidClassroomId error will be returned
// If passed comment is invalid, ErrInvalidComment error will be returned
func New(id uint64, userId uint64, classroomId uint64, comment string) (Feedback, error) {
	if id == 0 {
		return nil, ErrInvalidId
	}
	if userId == 0 {
		return nil, ErrInvalidUserId
	}
	if classroomId == 0 {
		return nil, ErrInvalidClassroomId
	}
	if len(comment) == 0 {
		return nil, ErrInvalidComment
	}
	return &feedback{id, userId, classroomId, comment}, nil
}

// Id returns feedback identifier
func (f *feedback) Id() uint64 {
	return f.id
}

// UserId returns user identifier
func (f *feedback) UserId() uint64 {
	return f.userId
}

// ClassroomId returns classroom identifier
func (f *feedback) ClassroomId() uint64 {
	return f.classroomId
}

// UpdateComment sets a new comment
// If passed comment is invalid, ErrInvalidComment error will be returned
func (f *feedback) UpdateComment(comment string) error {
	if len(comment) == 0 {
		return ErrInvalidComment
	}
	f.comment = comment
	return nil
}

// Comment returns a comment left by user
func (f *feedback) Comment() string {
	return f.comment
}

func (f *feedback) String() string {
	return fmt.Sprintf(formatString, f.id, f.userId, f.classroomId, f.comment)
}
