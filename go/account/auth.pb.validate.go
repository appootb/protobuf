// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: auth.proto

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
var _auth_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LoginRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RelateId

	switch m.Secure.(type) {

	case *LoginRequest_Code:
		// no validation rules for Code

	case *LoginRequest_Password:
		// no validation rules for Password

	}

	return nil
}

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on OAuthRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OAuthRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	// no validation rules for OpenId

	// no validation rules for AccessToken

	return nil
}

// OAuthRequestValidationError is the validation error returned by
// OAuthRequest.Validate if the designated constraints aren't met.
type OAuthRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuthRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuthRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuthRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuthRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuthRequestValidationError) ErrorName() string { return "OAuthRequestValidationError" }

// Error satisfies the builtin error interface
func (e OAuthRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuthRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuthRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuthRequestValidationError{}

// Validate checks the field values on Region with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Region) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Local

	// no validation rules for Abbr

	// no validation rules for Code

	return nil
}

// RegionValidationError is the validation error returned by Region.Validate if
// the designated constraints aren't met.
type RegionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegionValidationError) ErrorName() string { return "RegionValidationError" }

// Error satisfies the builtin error interface
func (e RegionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegionValidationError{}

// Validate checks the field values on Regions with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Regions) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRegions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegionsValidationError{
					field:  fmt.Sprintf("Regions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RegionsValidationError is the validation error returned by Regions.Validate
// if the designated constraints aren't met.
type RegionsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegionsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegionsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegionsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegionsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegionsValidationError) ErrorName() string { return "RegionsValidationError" }

// Error satisfies the builtin error interface
func (e RegionsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegions.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegionsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegionsValidationError{}