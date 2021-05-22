package feedback

import (
	"fmt"
)

const (
	ErrInvalidId      = FeedbackErr("provided id is invalid")
	ErrInvalidUserId  = FeedbackErr("provided user id is invalid")
	ErrInvalidComment = FeedbackErr("provided comment is invalid")

	formatString = "id: %v, user: %v, comment: %v"
)

// FeedbackErr are errors that happen while interacting with a feedback
type FeedbackErr string

func (e FeedbackErr) Error() string {
	return string(e)
}

type Feedback interface {
	Id() uint64
	UserId() uint64
	Comment() string
}

type feedback struct {
	id      uint64
	userId  uint64
	comment string
}

// New returns a Feedback object representing a user feedback
// If passed id is invalid, ErrInvalidId error will be returned
// If passed userId is invalid, ErrInvalidUserId error will be returned
// If passed comment is invalid, ErrInvalidComment error will be returned
func New(id uint64, userId uint64, comment string) (Feedback, error) {
	if id == 0 {
		return nil, ErrInvalidId
	}
	if userId == 0 {
		return nil, ErrInvalidUserId
	}
	if len(comment) == 0 {
		return nil, ErrInvalidComment
	}
	return &feedback{id, userId, comment}, nil
}

// Id returns feedback identifier
func (f *feedback) Id() uint64 {
	return f.id
}

// UserId returns user identifier
func (f *feedback) UserId() uint64 {
	return f.userId
}

// Comment returns a comment left by user
func (f *feedback) Comment() string {
	return f.comment
}

func (f *feedback) String() string {
	return fmt.Sprintf(formatString, f.id, f.userId, f.comment)
}
