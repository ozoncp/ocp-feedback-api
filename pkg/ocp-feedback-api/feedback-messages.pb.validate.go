// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: feedback-messages.proto

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

	if m.GetFeedbackId() <= 0 {
		return FeedbackValidationError{
			field:  "FeedbackId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetUserId() <= 0 {
		return FeedbackValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetClassroomId() <= 0 {
		return FeedbackValidationError{
			field:  "ClassroomId",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetComment()) < 1 {
		return FeedbackValidationError{
			field:  "Comment",
			reason: "value length must be at least 1 runes",
		}
	}

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

// Validate checks the field values on NewFeedback with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *NewFeedback) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUserId() <= 0 {
		return NewFeedbackValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetClassroomId() <= 0 {
		return NewFeedbackValidationError{
			field:  "ClassroomId",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetComment()) < 1 {
		return NewFeedbackValidationError{
			field:  "Comment",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// NewFeedbackValidationError is the validation error returned by
// NewFeedback.Validate if the designated constraints aren't met.
type NewFeedbackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewFeedbackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewFeedbackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewFeedbackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewFeedbackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewFeedbackValidationError) ErrorName() string { return "NewFeedbackValidationError" }

// Error satisfies the builtin error interface
func (e NewFeedbackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewFeedback.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewFeedbackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewFeedbackValidationError{}

// Validate checks the field values on CreateFeedbackV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateFeedbackV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetNewFeedback()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateFeedbackV1RequestValidationError{
				field:  "NewFeedback",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

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

	if m.GetFeedbackId() <= 0 {
		return CreateFeedbackV1ResponseValidationError{
			field:  "FeedbackId",
			reason: "value must be greater than 0",
		}
	}

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

// Validate checks the field values on CreateMultiFeedbackV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateMultiFeedbackV1Request) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetNewFeedbacks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateMultiFeedbackV1RequestValidationError{
					field:  fmt.Sprintf("NewFeedbacks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateMultiFeedbackV1RequestValidationError is the validation error returned
// by CreateMultiFeedbackV1Request.Validate if the designated constraints
// aren't met.
type CreateMultiFeedbackV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMultiFeedbackV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMultiFeedbackV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMultiFeedbackV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMultiFeedbackV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMultiFeedbackV1RequestValidationError) ErrorName() string {
	return "CreateMultiFeedbackV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMultiFeedbackV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMultiFeedbackV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMultiFeedbackV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMultiFeedbackV1RequestValidationError{}

// Validate checks the field values on CreateMultiFeedbackV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateMultiFeedbackV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CreateMultiFeedbackV1ResponseValidationError is the validation error
// returned by CreateMultiFeedbackV1Response.Validate if the designated
// constraints aren't met.
type CreateMultiFeedbackV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMultiFeedbackV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMultiFeedbackV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMultiFeedbackV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMultiFeedbackV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMultiFeedbackV1ResponseValidationError) ErrorName() string {
	return "CreateMultiFeedbackV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMultiFeedbackV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMultiFeedbackV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMultiFeedbackV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMultiFeedbackV1ResponseValidationError{}

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
