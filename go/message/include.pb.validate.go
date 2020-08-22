// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: include.proto

package message

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

// Validate checks the field values on Postmark with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Postmark) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Category

	// no validation rules for Sender

	// no validation rules for Receiver

	// no validation rules for UniqueId

	if v, ok := interface{}(m.GetExpire()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PostmarkValidationError{
				field:  "Expire",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PostmarkValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PostmarkValidationError is the validation error returned by
// Postmark.Validate if the designated constraints aren't met.
type PostmarkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostmarkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostmarkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostmarkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostmarkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostmarkValidationError) ErrorName() string { return "PostmarkValidationError" }

// Error satisfies the builtin error interface
func (e PostmarkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPostmark.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostmarkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostmarkValidationError{}

// Validate checks the field values on Text with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Text) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Content

	return nil
}

// TextValidationError is the validation error returned by Text.Validate if the
// designated constraints aren't met.
type TextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TextValidationError) ErrorName() string { return "TextValidationError" }

// Error satisfies the builtin error interface
func (e TextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sText.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TextValidationError{}

// Validate checks the field values on Media with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Media) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Url

	// no validation rules for Thumbnail

	return nil
}

// MediaValidationError is the validation error returned by Media.Validate if
// the designated constraints aren't met.
type MediaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MediaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MediaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MediaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MediaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MediaValidationError) ErrorName() string { return "MediaValidationError" }

// Error satisfies the builtin error interface
func (e MediaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMedia.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MediaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MediaValidationError{}

// Validate checks the field values on Location with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Location) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Latitude

	// no validation rules for Longitude

	return nil
}

// LocationValidationError is the validation error returned by
// Location.Validate if the designated constraints aren't met.
type LocationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocationValidationError) ErrorName() string { return "LocationValidationError" }

// Error satisfies the builtin error interface
func (e LocationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocationValidationError{}

// Validate checks the field values on Post with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Post) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetPostmark()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PostValidationError{
				field:  "Postmark",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Type

	switch m.Envelope.(type) {

	case *Post_Text:

		if v, ok := interface{}(m.GetText()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PostValidationError{
					field:  "Text",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Post_Resource:

		if v, ok := interface{}(m.GetResource()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PostValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Post_Coordinate:

		if v, ok := interface{}(m.GetCoordinate()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PostValidationError{
					field:  "Coordinate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PostValidationError is the validation error returned by Post.Validate if the
// designated constraints aren't met.
type PostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostValidationError) ErrorName() string { return "PostValidationError" }

// Error satisfies the builtin error interface
func (e PostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostValidationError{}

// Validate checks the field values on Postbox with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Postbox) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PostboxValidationError{
					field:  fmt.Sprintf("Posts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PostboxValidationError is the validation error returned by Postbox.Validate
// if the designated constraints aren't met.
type PostboxValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PostboxValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PostboxValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PostboxValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PostboxValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PostboxValidationError) ErrorName() string { return "PostboxValidationError" }

// Error satisfies the builtin error interface
func (e PostboxValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPostbox.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PostboxValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PostboxValidationError{}

// Validate checks the field values on Receipts with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Receipts) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetPostmarks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReceiptsValidationError{
					field:  fmt.Sprintf("Postmarks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ReceiptsValidationError is the validation error returned by
// Receipts.Validate if the designated constraints aren't met.
type ReceiptsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReceiptsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReceiptsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReceiptsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReceiptsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReceiptsValidationError) ErrorName() string { return "ReceiptsValidationError" }

// Error satisfies the builtin error interface
func (e ReceiptsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReceipts.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReceiptsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReceiptsValidationError{}

// Validate checks the field values on Stream with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Stream) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Sn

	// no validation rules for Type

	switch m.Payload.(type) {

	case *Stream_Message:

		if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Stream_Postmark:

		if v, ok := interface{}(m.GetPostmark()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamValidationError{
					field:  "Postmark",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Stream_Posts:

		if v, ok := interface{}(m.GetPosts()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamValidationError{
					field:  "Posts",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Stream_Receipts:

		if v, ok := interface{}(m.GetReceipts()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StreamValidationError{
					field:  "Receipts",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// StreamValidationError is the validation error returned by Stream.Validate if
// the designated constraints aren't met.
type StreamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamValidationError) ErrorName() string { return "StreamValidationError" }

// Error satisfies the builtin error interface
func (e StreamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStream.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamValidationError{}

// Validate checks the field values on SessionCast with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SessionCast) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Product

	return nil
}

// SessionCastValidationError is the validation error returned by
// SessionCast.Validate if the designated constraints aren't met.
type SessionCastValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SessionCastValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SessionCastValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SessionCastValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SessionCastValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SessionCastValidationError) ErrorName() string { return "SessionCastValidationError" }

// Error satisfies the builtin error interface
func (e SessionCastValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSessionCast.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SessionCastValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SessionCastValidationError{}

// Validate checks the field values on TagBroadcast with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *TagBroadcast) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Prefix

	return nil
}

// TagBroadcastValidationError is the validation error returned by
// TagBroadcast.Validate if the designated constraints aren't met.
type TagBroadcastValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TagBroadcastValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TagBroadcastValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TagBroadcastValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TagBroadcastValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TagBroadcastValidationError) ErrorName() string { return "TagBroadcastValidationError" }

// Error satisfies the builtin error interface
func (e TagBroadcastValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTagBroadcast.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TagBroadcastValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TagBroadcastValidationError{}

// Validate checks the field values on Package with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Package) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Sn

	for idx, item := range m.GetPosts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PackageValidationError{
					field:  fmt.Sprintf("Posts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.Target.(type) {

	case *Package_Session:

		if v, ok := interface{}(m.GetSession()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PackageValidationError{
					field:  "Session",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Package_Tag:

		if v, ok := interface{}(m.GetTag()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PackageValidationError{
					field:  "Tag",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// PackageValidationError is the validation error returned by Package.Validate
// if the designated constraints aren't met.
type PackageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PackageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PackageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PackageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PackageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PackageValidationError) ErrorName() string { return "PackageValidationError" }

// Error satisfies the builtin error interface
func (e PackageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPackage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PackageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PackageValidationError{}