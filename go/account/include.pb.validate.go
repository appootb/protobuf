// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: include.proto

package account

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
var _include_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Secret with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Secret) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	return nil
}

// SecretValidationError is the validation error returned by Secret.Validate if
// the designated constraints aren't met.
type SecretValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SecretValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SecretValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SecretValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SecretValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SecretValidationError) ErrorName() string { return "SecretValidationError" }

// Error satisfies the builtin error interface
func (e SecretValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSecret.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SecretValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SecretValidationError{}

// Validate checks the field values on Info with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Info) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UniqueId

	// no validation rules for Nickname

	// no validation rules for Avatar

	// no validation rules for Status

	// no validation rules for Signature

	// no validation rules for Gender

	// no validation rules for Signs

	// no validation rules for Location

	if v, ok := interface{}(m.GetSecret()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InfoValidationError{
				field:  "Secret",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InfoValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InfoValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetExtend()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InfoValidationError{
				field:  "Extend",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// InfoValidationError is the validation error returned by Info.Validate if the
// designated constraints aren't met.
type InfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfoValidationError) ErrorName() string { return "InfoValidationError" }

// Error satisfies the builtin error interface
func (e InfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfoValidationError{}

// Validate checks the field values on Infos with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Infos) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetAccounts() {
		_ = val

		// no validation rules for Accounts[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return InfosValidationError{
					field:  fmt.Sprintf("Accounts[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// InfosValidationError is the validation error returned by Infos.Validate if
// the designated constraints aren't met.
type InfosValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InfosValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InfosValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InfosValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InfosValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InfosValidationError) ErrorName() string { return "InfosValidationError" }

// Error satisfies the builtin error interface
func (e InfosValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInfos.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InfosValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InfosValidationError{}
