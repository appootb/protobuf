// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common.proto

package captcha

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
var _common_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CodeRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CodeRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Channel

	// no validation rules for Category

	// no validation rules for Target

	// no validation rules for Value

	return nil
}

// CodeRequestValidationError is the validation error returned by
// CodeRequest.Validate if the designated constraints aren't met.
type CodeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CodeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CodeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CodeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CodeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CodeRequestValidationError) ErrorName() string { return "CodeRequestValidationError" }

// Error satisfies the builtin error interface
func (e CodeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCodeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CodeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CodeRequestValidationError{}
