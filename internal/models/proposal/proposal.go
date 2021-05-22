package proposal

import (
	"fmt"

	"github.com/ozoncp/ocp-feedback-api/internal/models/entity"
)

const (
	ErrInvalidId         = ProposalErr("provided propsal id is invalid")
	ErrInvalidUserId     = ProposalErr("provided user id is invalid")
	ErrInvalidLessonId   = ProposalErr("provided lesson id is invalid")
	ErrInvalidDocumentId = ProposalErr("provided document id is invalid")

	formatString = "id: %v, user: %v, lesson: %v, document: %v"
)

// ProposalErr are errors that happen while interacting with a feedback
type ProposalErr string

func (e ProposalErr) Error() string {
	return string(e)
}

type Proposal interface {
	entity.Entity
	LessonId() uint64
	DocumentId() uint64
}

type proposal struct {
	id         uint64
	userId     uint64
	lessonId   uint64
	documentId uint64
}

// New returns a Proposal object representing a user proposal
// If passed id is invalid, ErrInvalidId error will be returned
// If passed user id is invalid, ErrInvalidUserId error will be returned
// If passed lessonId is invalid, ErrInvalidLessonId error will be returned
// If passed documentId is invalid, ErrInvalidDocumentId error will be returned
func New(id, userId, lessonId, documentId uint64) (Proposal, error) {
	if id == 0 {
		return nil, ErrInvalidId
	}
	if userId == 0 {
		return nil, ErrInvalidUserId
	}
	if lessonId == 0 {
		return nil, ErrInvalidLessonId
	}
	if documentId == 0 {
		return nil, ErrInvalidDocumentId
	}
	return &proposal{id, userId, lessonId, documentId}, nil
}

// Id returns a proposal id
func (p *proposal) Id() uint64 {
	return p.id
}

// UserId returns an id of the user who made this proposal
func (p *proposal) UserId() uint64 {
	return p.userId
}

// LessonId returns a lesson id mentioned in the proposal
func (p *proposal) LessonId() uint64 {
	return p.lessonId
}

// DocumentId returns a document id mentioned in the proposal
func (p *proposal) DocumentId() uint64 {
	return p.documentId
}

func (p *proposal) String() string {
	return fmt.Sprintf(formatString, p.id, p.userId, p.lessonId, p.documentId)
}
