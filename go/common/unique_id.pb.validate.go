// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: unique_id.proto

package common

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _unique_id_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on UniqueId with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UniqueId) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// UniqueIdValidationError is the validation error returned by
// UniqueId.Validate if the designated constraints aren't met.
type UniqueIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UniqueIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UniqueIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UniqueIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UniqueIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UniqueIdValidationError) ErrorName() string { return "UniqueIdValidationError" }

// Error satisfies the builtin error interface
func (e UniqueIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUniqueId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UniqueIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UniqueIdValidationError{}

// Validate checks the field values on UniqueIds with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UniqueIds) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UniqueIdsValidationError is the validation error returned by
// UniqueIds.Validate if the designated constraints aren't met.
type UniqueIdsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UniqueIdsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UniqueIdsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UniqueIdsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UniqueIdsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UniqueIdsValidationError) ErrorName() string { return "UniqueIdsValidationError" }

// Error satisfies the builtin error interface
func (e UniqueIdsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUniqueIds.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UniqueIdsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UniqueIdsValidationError{}