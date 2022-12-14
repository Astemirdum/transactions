// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: balance/v1/balance.proto

package balancev1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on CreateBalanceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateBalanceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateBalanceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateBalanceRequestMultiError, or nil if none found.
func (m *CreateBalanceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateBalanceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetUserId() < 0 {
		err := CreateBalanceRequestValidationError{
			field:  "UserId",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateBalanceRequestMultiError(errors)
	}

	return nil
}

// CreateBalanceRequestMultiError is an error wrapping multiple validation
// errors returned by CreateBalanceRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateBalanceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateBalanceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateBalanceRequestMultiError) AllErrors() []error { return m }

// CreateBalanceRequestValidationError is the validation error returned by
// CreateBalanceRequest.Validate if the designated constraints aren't met.
type CreateBalanceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateBalanceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateBalanceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateBalanceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateBalanceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateBalanceRequestValidationError) ErrorName() string {
	return "CreateBalanceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateBalanceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateBalanceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateBalanceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateBalanceRequestValidationError{}

// Validate checks the field values on CreateBalanceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateBalanceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateBalanceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateBalanceResponseMultiError, or nil if none found.
func (m *CreateBalanceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateBalanceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Cash

	if len(errors) > 0 {
		return CreateBalanceResponseMultiError(errors)
	}

	return nil
}

// CreateBalanceResponseMultiError is an error wrapping multiple validation
// errors returned by CreateBalanceResponse.ValidateAll() if the designated
// constraints aren't met.
type CreateBalanceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateBalanceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateBalanceResponseMultiError) AllErrors() []error { return m }

// CreateBalanceResponseValidationError is the validation error returned by
// CreateBalanceResponse.Validate if the designated constraints aren't met.
type CreateBalanceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateBalanceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateBalanceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateBalanceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateBalanceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateBalanceResponseValidationError) ErrorName() string {
	return "CreateBalanceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateBalanceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateBalanceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateBalanceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateBalanceResponseValidationError{}

// Validate checks the field values on CashOutRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CashOutRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CashOutRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CashOutRequestMultiError,
// or nil if none found.
func (m *CashOutRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CashOutRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCash() <= 0 {
		err := CashOutRequestValidationError{
			field:  "Cash",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CashOutRequestMultiError(errors)
	}

	return nil
}

// CashOutRequestMultiError is an error wrapping multiple validation errors
// returned by CashOutRequest.ValidateAll() if the designated constraints
// aren't met.
type CashOutRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CashOutRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CashOutRequestMultiError) AllErrors() []error { return m }

// CashOutRequestValidationError is the validation error returned by
// CashOutRequest.Validate if the designated constraints aren't met.
type CashOutRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CashOutRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CashOutRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CashOutRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CashOutRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CashOutRequestValidationError) ErrorName() string { return "CashOutRequestValidationError" }

// Error satisfies the builtin error interface
func (e CashOutRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCashOutRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CashOutRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CashOutRequestValidationError{}

// Validate checks the field values on CashOutResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CashOutResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CashOutResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CashOutResponseMultiError, or nil if none found.
func (m *CashOutResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CashOutResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CashOutResponseMultiError(errors)
	}

	return nil
}

// CashOutResponseMultiError is an error wrapping multiple validation errors
// returned by CashOutResponse.ValidateAll() if the designated constraints
// aren't met.
type CashOutResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CashOutResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CashOutResponseMultiError) AllErrors() []error { return m }

// CashOutResponseValidationError is the validation error returned by
// CashOutResponse.Validate if the designated constraints aren't met.
type CashOutResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CashOutResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CashOutResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CashOutResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CashOutResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CashOutResponseValidationError) ErrorName() string { return "CashOutResponseValidationError" }

// Error satisfies the builtin error interface
func (e CashOutResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCashOutResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CashOutResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CashOutResponseValidationError{}

// Validate checks the field values on GetBalanceRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetBalanceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBalanceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetBalanceRequestMultiError, or nil if none found.
func (m *GetBalanceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBalanceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetBalanceRequestMultiError(errors)
	}

	return nil
}

// GetBalanceRequestMultiError is an error wrapping multiple validation errors
// returned by GetBalanceRequest.ValidateAll() if the designated constraints
// aren't met.
type GetBalanceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBalanceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBalanceRequestMultiError) AllErrors() []error { return m }

// GetBalanceRequestValidationError is the validation error returned by
// GetBalanceRequest.Validate if the designated constraints aren't met.
type GetBalanceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBalanceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBalanceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBalanceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBalanceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBalanceRequestValidationError) ErrorName() string {
	return "GetBalanceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetBalanceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBalanceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBalanceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBalanceRequestValidationError{}

// Validate checks the field values on GetBalanceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetBalanceResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetBalanceResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetBalanceResponseMultiError, or nil if none found.
func (m *GetBalanceResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetBalanceResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Cash

	if len(errors) > 0 {
		return GetBalanceResponseMultiError(errors)
	}

	return nil
}

// GetBalanceResponseMultiError is an error wrapping multiple validation errors
// returned by GetBalanceResponse.ValidateAll() if the designated constraints
// aren't met.
type GetBalanceResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetBalanceResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetBalanceResponseMultiError) AllErrors() []error { return m }

// GetBalanceResponseValidationError is the validation error returned by
// GetBalanceResponse.Validate if the designated constraints aren't met.
type GetBalanceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetBalanceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetBalanceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetBalanceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetBalanceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetBalanceResponseValidationError) ErrorName() string {
	return "GetBalanceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetBalanceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetBalanceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetBalanceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetBalanceResponseValidationError{}
