// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: inner_key.proto

package counter

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
var _inner_key_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Key with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Key) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Product

	// no validation rules for Type

	// no validation rules for RelateId

	// no validation rules for Value

	if v, ok := interface{}(m.GetExpire()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KeyValidationError{
				field:  "Expire",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// KeyValidationError is the validation error returned by Key.Validate if the
// designated constraints aren't met.
type KeyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyValidationError) ErrorName() string { return "KeyValidationError" }

// Error satisfies the builtin error interface
func (e KeyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKey.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyValidationError{}

// Validate checks the field values on Keys with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Keys) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Product

	// no validation rules for Type

	// no validation rules for Values

	return nil
}

// KeysValidationError is the validation error returned by Keys.Validate if the
// designated constraints aren't met.
type KeysValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeysValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeysValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeysValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeysValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeysValidationError) ErrorName() string { return "KeysValidationError" }

// Error satisfies the builtin error interface
func (e KeysValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeys.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeysValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeysValidationError{}
