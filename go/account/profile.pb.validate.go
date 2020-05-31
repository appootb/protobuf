// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: profile.proto

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
var _profile_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AccountProfile with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *AccountProfile) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AccountProfileValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AccountProfileValidationError is the validation error returned by
// AccountProfile.Validate if the designated constraints aren't met.
type AccountProfileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountProfileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountProfileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountProfileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountProfileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountProfileValidationError) ErrorName() string { return "AccountProfileValidationError" }

// Error satisfies the builtin error interface
func (e AccountProfileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountProfile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountProfileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountProfileValidationError{}

// Validate checks the field values on AccountProfiles with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AccountProfiles) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetProfiles() {
		_ = val

		// no validation rules for Profiles[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AccountProfilesValidationError{
					field:  fmt.Sprintf("Profiles[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AccountProfilesValidationError is the validation error returned by
// AccountProfiles.Validate if the designated constraints aren't met.
type AccountProfilesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AccountProfilesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AccountProfilesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AccountProfilesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AccountProfilesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AccountProfilesValidationError) ErrorName() string { return "AccountProfilesValidationError" }

// Error satisfies the builtin error interface
func (e AccountProfilesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAccountProfiles.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AccountProfilesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AccountProfilesValidationError{}
