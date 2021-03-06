// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proposal-messages.proto

package ocp_proposal_api

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

// Validate checks the field values on Proposal with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Proposal) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ProposalId

	if m.GetUserId() <= 0 {
		return ProposalValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetLessonId() <= 0 {
		return ProposalValidationError{
			field:  "LessonId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetDocumentId() <= 0 {
		return ProposalValidationError{
			field:  "DocumentId",
			reason: "value must be greater than 0",
		}
	}

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

	if m.GetProposal() == nil {
		return CreateProposalV1RequestValidationError{
			field:  "Proposal",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetProposal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateProposalV1RequestValidationError{
				field:  "Proposal",
				reason: "embedded message failed validation",
				cause:  err,
			}
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

	if m.GetProposalId() <= 0 {
		return CreateProposalV1ResponseValidationError{
			field:  "ProposalId",
			reason: "value must be greater than 0",
		}
	}

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

// Validate checks the field values on CreateMultiProposalV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateMultiProposalV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetProposals()) < 1 {
		return CreateMultiProposalV1RequestValidationError{
			field:  "Proposals",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetProposals() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateMultiProposalV1RequestValidationError{
					field:  fmt.Sprintf("Proposals[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CreateMultiProposalV1RequestValidationError is the validation error returned
// by CreateMultiProposalV1Request.Validate if the designated constraints
// aren't met.
type CreateMultiProposalV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMultiProposalV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMultiProposalV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMultiProposalV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMultiProposalV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMultiProposalV1RequestValidationError) ErrorName() string {
	return "CreateMultiProposalV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMultiProposalV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMultiProposalV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMultiProposalV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMultiProposalV1RequestValidationError{}

// Validate checks the field values on CreateMultiProposalV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateMultiProposalV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetProposals()) < 1 {
		return CreateMultiProposalV1ResponseValidationError{
			field:  "Proposals",
			reason: "value must contain at least 1 item(s)",
		}
	}

	return nil
}

// CreateMultiProposalV1ResponseValidationError is the validation error
// returned by CreateMultiProposalV1Response.Validate if the designated
// constraints aren't met.
type CreateMultiProposalV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateMultiProposalV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateMultiProposalV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateMultiProposalV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateMultiProposalV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateMultiProposalV1ResponseValidationError) ErrorName() string {
	return "CreateMultiProposalV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateMultiProposalV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateMultiProposalV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateMultiProposalV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateMultiProposalV1ResponseValidationError{}

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

	if m.GetProposal() == nil {
		return DescribeProposalV1ResponseValidationError{
			field:  "Proposal",
			reason: "value is required",
		}
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

// Validate checks the field values on UpdateProposalV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateProposalV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetProposal() == nil {
		return UpdateProposalV1RequestValidationError{
			field:  "Proposal",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetProposal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateProposalV1RequestValidationError{
				field:  "Proposal",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateProposalV1RequestValidationError is the validation error returned by
// UpdateProposalV1Request.Validate if the designated constraints aren't met.
type UpdateProposalV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProposalV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProposalV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProposalV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProposalV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProposalV1RequestValidationError) ErrorName() string {
	return "UpdateProposalV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProposalV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProposalV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProposalV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProposalV1RequestValidationError{}

// Validate checks the field values on UpdateProposalV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateProposalV1Response) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateProposalV1ResponseValidationError is the validation error returned by
// UpdateProposalV1Response.Validate if the designated constraints aren't met.
type UpdateProposalV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateProposalV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateProposalV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateProposalV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateProposalV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateProposalV1ResponseValidationError) ErrorName() string {
	return "UpdateProposalV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateProposalV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateProposalV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateProposalV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateProposalV1ResponseValidationError{}
