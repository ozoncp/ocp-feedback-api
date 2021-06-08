// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-feedback-api/ocp-feedback-api.proto

package ocp_feedback_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on Feedback with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Feedback) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FeedbackId

	// no validation rules for UserId

	// no validation rules for ClassroomId

	// no validation rules for Comment

	return nil
}

// FeedbackValidationError is the validation error returned by
// Feedback.Validate if the designated constraints aren't met.
type FeedbackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeedbackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeedbackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeedbackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeedbackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeedbackValidationError) ErrorName() string { return "FeedbackValidationError" }

// Error satisfies the builtin error interface
func (e FeedbackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeedback.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeedbackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeedbackValidationError{}

// Validate checks the field values on CreateFeedbackV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateFeedbackV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return CreateFeedbackV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetClassroomId() <= 0 {
		return CreateFeedbackV1RequestValidationError{
			field:  "ClassroomId",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Comment

	return nil
}

// CreateFeedbackV1RequestValidationError is the validation error returned by
// CreateFeedbackV1Request.Validate if the designated constraints aren't met.
type CreateFeedbackV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateFeedbackV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateFeedbackV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateFeedbackV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateFeedbackV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateFeedbackV1RequestValidationError) ErrorName() string {
	return "CreateFeedbackV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateFeedbackV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateFeedbackV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateFeedbackV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateFeedbackV1RequestValidationError{}

// Validate checks the field values on CreateFeedbackV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateFeedbackV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FeedbackId

	return nil
}

// CreateFeedbackV1ResponseValidationError is the validation error returned by
// CreateFeedbackV1Response.Validate if the designated constraints aren't met.
type CreateFeedbackV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateFeedbackV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateFeedbackV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateFeedbackV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateFeedbackV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateFeedbackV1ResponseValidationError) ErrorName() string {
	return "CreateFeedbackV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateFeedbackV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateFeedbackV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateFeedbackV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateFeedbackV1ResponseValidationError{}

// Validate checks the field values on RemoveFeedbackV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveFeedbackV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetFeedbackId() <= 0 {
		return RemoveFeedbackV1RequestValidationError{
			field:  "FeedbackId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveFeedbackV1RequestValidationError is the validation error returned by
// RemoveFeedbackV1Request.Validate if the designated constraints aren't met.
type RemoveFeedbackV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveFeedbackV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveFeedbackV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveFeedbackV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveFeedbackV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveFeedbackV1RequestValidationError) ErrorName() string {
	return "RemoveFeedbackV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveFeedbackV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveFeedbackV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveFeedbackV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveFeedbackV1RequestValidationError{}

// Validate checks the field values on RemoveFeedbackV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveFeedbackV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RemoveFeedbackV1ResponseValidationError is the validation error returned by
// RemoveFeedbackV1Response.Validate if the designated constraints aren't met.
type RemoveFeedbackV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveFeedbackV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveFeedbackV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveFeedbackV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveFeedbackV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveFeedbackV1ResponseValidationError) ErrorName() string {
	return "RemoveFeedbackV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveFeedbackV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveFeedbackV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveFeedbackV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveFeedbackV1ResponseValidationError{}

// Validate checks the field values on DescribeFeedbackV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeFeedbackV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetFeedbackId() <= 0 {
		return DescribeFeedbackV1RequestValidationError{
			field:  "FeedbackId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeFeedbackV1RequestValidationError is the validation error returned by
// DescribeFeedbackV1Request.Validate if the designated constraints aren't met.
type DescribeFeedbackV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeFeedbackV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeFeedbackV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeFeedbackV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeFeedbackV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeFeedbackV1RequestValidationError) ErrorName() string {
	return "DescribeFeedbackV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeFeedbackV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeFeedbackV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeFeedbackV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeFeedbackV1RequestValidationError{}

// Validate checks the field values on DescribeFeedbackV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeFeedbackV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetFeedback()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeFeedbackV1ResponseValidationError{
				field:  "Feedback",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeFeedbackV1ResponseValidationError is the validation error returned
// by DescribeFeedbackV1Response.Validate if the designated constraints aren't met.
type DescribeFeedbackV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeFeedbackV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeFeedbackV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeFeedbackV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeFeedbackV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeFeedbackV1ResponseValidationError) ErrorName() string {
	return "DescribeFeedbackV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeFeedbackV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeFeedbackV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeFeedbackV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeFeedbackV1ResponseValidationError{}

// Validate checks the field values on ListFeedbacksV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListFeedbacksV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimit() <= 0 {
		return ListFeedbacksV1RequestValidationError{
			field:  "Limit",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Offset

	return nil
}

// ListFeedbacksV1RequestValidationError is the validation error returned by
// ListFeedbacksV1Request.Validate if the designated constraints aren't met.
type ListFeedbacksV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFeedbacksV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFeedbacksV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFeedbacksV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFeedbacksV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFeedbacksV1RequestValidationError) ErrorName() string {
	return "ListFeedbacksV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListFeedbacksV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFeedbacksV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFeedbacksV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFeedbacksV1RequestValidationError{}

// Validate checks the field values on ListFeedbacksV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListFeedbacksV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetFeedbacks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListFeedbacksV1ResponseValidationError{
					field:  fmt.Sprintf("Feedbacks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListFeedbacksV1ResponseValidationError is the validation error returned by
// ListFeedbacksV1Response.Validate if the designated constraints aren't met.
type ListFeedbacksV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListFeedbacksV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListFeedbacksV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListFeedbacksV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListFeedbacksV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListFeedbacksV1ResponseValidationError) ErrorName() string {
	return "ListFeedbacksV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListFeedbacksV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListFeedbacksV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListFeedbacksV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListFeedbacksV1ResponseValidationError{}

// Validate checks the field values on Proposal with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Proposal) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProposalId

	// no validation rules for UserId

	// no validation rules for LessonId

	// no validation rules for DocumentId

	return nil
}

// ProposalValidationError is the validation error returned by
// Proposal.Validate if the designated constraints aren't met.
type ProposalValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProposalValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProposalValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProposalValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProposalValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProposalValidationError) ErrorName() string { return "ProposalValidationError" }

// Error satisfies the builtin error interface
func (e ProposalValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProposal.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProposalValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProposalValidationError{}

// Validate checks the field values on CreateProposalV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateProposalV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return CreateProposalV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetLessonId() <= 0 {
		return CreateProposalV1RequestValidationError{
			field:  "LessonId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetDocumentId() <= 0 {
		return CreateProposalV1RequestValidationError{
			field:  "DocumentId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CreateProposalV1RequestValidationError is the validation error returned by
// CreateProposalV1Request.Validate if the designated constraints aren't met.
type CreateProposalV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProposalV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProposalV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProposalV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProposalV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProposalV1RequestValidationError) ErrorName() string {
	return "CreateProposalV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProposalV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProposalV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProposalV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProposalV1RequestValidationError{}

// Validate checks the field values on CreateProposalV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateProposalV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProposalId

	return nil
}

// CreateProposalV1ResponseValidationError is the validation error returned by
// CreateProposalV1Response.Validate if the designated constraints aren't met.
type CreateProposalV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProposalV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProposalV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProposalV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProposalV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProposalV1ResponseValidationError) ErrorName() string {
	return "CreateProposalV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProposalV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProposalV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProposalV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProposalV1ResponseValidationError{}

// Validate checks the field values on RemoveProposalV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveProposalV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetProposalId() <= 0 {
		return RemoveProposalV1RequestValidationError{
			field:  "ProposalId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveProposalV1RequestValidationError is the validation error returned by
// RemoveProposalV1Request.Validate if the designated constraints aren't met.
type RemoveProposalV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveProposalV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveProposalV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveProposalV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveProposalV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveProposalV1RequestValidationError) ErrorName() string {
	return "RemoveProposalV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveProposalV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveProposalV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveProposalV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveProposalV1RequestValidationError{}

// Validate checks the field values on RemoveProposalV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveProposalV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RemoveProposalV1ResponseValidationError is the validation error returned by
// RemoveProposalV1Response.Validate if the designated constraints aren't met.
type RemoveProposalV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveProposalV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveProposalV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveProposalV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveProposalV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveProposalV1ResponseValidationError) ErrorName() string {
	return "RemoveProposalV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveProposalV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveProposalV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveProposalV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveProposalV1ResponseValidationError{}

// Validate checks the field values on DescribeProposalV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeProposalV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetProposalId() <= 0 {
		return DescribeProposalV1RequestValidationError{
			field:  "ProposalId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeProposalV1RequestValidationError is the validation error returned by
// DescribeProposalV1Request.Validate if the designated constraints aren't met.
type DescribeProposalV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeProposalV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeProposalV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeProposalV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeProposalV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeProposalV1RequestValidationError) ErrorName() string {
	return "DescribeProposalV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeProposalV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeProposalV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeProposalV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeProposalV1RequestValidationError{}

// Validate checks the field values on DescribeProposalV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeProposalV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetProposal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeProposalV1ResponseValidationError{
				field:  "Proposal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeProposalV1ResponseValidationError is the validation error returned
// by DescribeProposalV1Response.Validate if the designated constraints aren't met.
type DescribeProposalV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeProposalV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeProposalV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeProposalV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeProposalV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeProposalV1ResponseValidationError) ErrorName() string {
	return "DescribeProposalV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeProposalV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeProposalV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeProposalV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeProposalV1ResponseValidationError{}

// Validate checks the field values on ListProposalsV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListProposalsV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLimit() <= 0 {
		return ListProposalsV1RequestValidationError{
			field:  "Limit",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for Offset

	return nil
}

// ListProposalsV1RequestValidationError is the validation error returned by
// ListProposalsV1Request.Validate if the designated constraints aren't met.
type ListProposalsV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProposalsV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProposalsV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProposalsV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProposalsV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProposalsV1RequestValidationError) ErrorName() string {
	return "ListProposalsV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListProposalsV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProposalsV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProposalsV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProposalsV1RequestValidationError{}

// Validate checks the field values on ListProposalsV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListProposalsV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetProposals() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListProposalsV1ResponseValidationError{
					field:  fmt.Sprintf("Proposals[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListProposalsV1ResponseValidationError is the validation error returned by
// ListProposalsV1Response.Validate if the designated constraints aren't met.
type ListProposalsV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListProposalsV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListProposalsV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListProposalsV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListProposalsV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListProposalsV1ResponseValidationError) ErrorName() string {
	return "ListProposalsV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListProposalsV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListProposalsV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListProposalsV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListProposalsV1ResponseValidationError{}
