// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: music/v1/music.proto

package v1

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

// Validate checks the field values on Track with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Track) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Track with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TrackMultiError, or nil if none found.
func (m *Track) ValidateAll() error {
	return m.validate(true)
}

func (m *Track) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Milliseconds

	// no validation rules for Bytes

	if len(errors) > 0 {
		return TrackMultiError(errors)
	}

	return nil
}

// TrackMultiError is an error wrapping multiple validation errors returned by
// Track.ValidateAll() if the designated constraints aren't met.
type TrackMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TrackMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TrackMultiError) AllErrors() []error { return m }

// TrackValidationError is the validation error returned by Track.Validate if
// the designated constraints aren't met.
type TrackValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TrackValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TrackValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TrackValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TrackValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TrackValidationError) ErrorName() string { return "TrackValidationError" }

// Error satisfies the builtin error interface
func (e TrackValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrack.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TrackValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TrackValidationError{}

// Validate checks the field values on AddTrackRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddTrackRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddTrackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddTrackRequestMultiError, or nil if none found.
func (m *AddTrackRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddTrackRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PlaylistId

	// no validation rules for TrackId

	if len(errors) > 0 {
		return AddTrackRequestMultiError(errors)
	}

	return nil
}

// AddTrackRequestMultiError is an error wrapping multiple validation errors
// returned by AddTrackRequest.ValidateAll() if the designated constraints
// aren't met.
type AddTrackRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddTrackRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddTrackRequestMultiError) AllErrors() []error { return m }

// AddTrackRequestValidationError is the validation error returned by
// AddTrackRequest.Validate if the designated constraints aren't met.
type AddTrackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddTrackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddTrackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddTrackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddTrackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddTrackRequestValidationError) ErrorName() string { return "AddTrackRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddTrackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddTrackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddTrackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddTrackRequestValidationError{}

// Validate checks the field values on AddTrackReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddTrackReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddTrackReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddTrackReplyMultiError, or
// nil if none found.
func (m *AddTrackReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AddTrackReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return AddTrackReplyMultiError(errors)
	}

	return nil
}

// AddTrackReplyMultiError is an error wrapping multiple validation errors
// returned by AddTrackReply.ValidateAll() if the designated constraints
// aren't met.
type AddTrackReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddTrackReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddTrackReplyMultiError) AllErrors() []error { return m }

// AddTrackReplyValidationError is the validation error returned by
// AddTrackReply.Validate if the designated constraints aren't met.
type AddTrackReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddTrackReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddTrackReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddTrackReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddTrackReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddTrackReplyValidationError) ErrorName() string { return "AddTrackReplyValidationError" }

// Error satisfies the builtin error interface
func (e AddTrackReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddTrackReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddTrackReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddTrackReplyValidationError{}

// Validate checks the field values on AddAlbumRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddAlbumRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddAlbumRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddAlbumRequestMultiError, or nil if none found.
func (m *AddAlbumRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddAlbumRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PlaylistId

	// no validation rules for AlbumId

	if len(errors) > 0 {
		return AddAlbumRequestMultiError(errors)
	}

	return nil
}

// AddAlbumRequestMultiError is an error wrapping multiple validation errors
// returned by AddAlbumRequest.ValidateAll() if the designated constraints
// aren't met.
type AddAlbumRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddAlbumRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddAlbumRequestMultiError) AllErrors() []error { return m }

// AddAlbumRequestValidationError is the validation error returned by
// AddAlbumRequest.Validate if the designated constraints aren't met.
type AddAlbumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddAlbumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddAlbumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddAlbumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddAlbumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddAlbumRequestValidationError) ErrorName() string { return "AddAlbumRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddAlbumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddAlbumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddAlbumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddAlbumRequestValidationError{}

// Validate checks the field values on AddAlbumReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddAlbumReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddAlbumReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddAlbumReplyMultiError, or
// nil if none found.
func (m *AddAlbumReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AddAlbumReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return AddAlbumReplyMultiError(errors)
	}

	return nil
}

// AddAlbumReplyMultiError is an error wrapping multiple validation errors
// returned by AddAlbumReply.ValidateAll() if the designated constraints
// aren't met.
type AddAlbumReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddAlbumReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddAlbumReplyMultiError) AllErrors() []error { return m }

// AddAlbumReplyValidationError is the validation error returned by
// AddAlbumReply.Validate if the designated constraints aren't met.
type AddAlbumReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddAlbumReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddAlbumReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddAlbumReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddAlbumReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddAlbumReplyValidationError) ErrorName() string { return "AddAlbumReplyValidationError" }

// Error satisfies the builtin error interface
func (e AddAlbumReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddAlbumReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddAlbumReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddAlbumReplyValidationError{}

// Validate checks the field values on DeleteTrackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteTrackRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTrackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTrackRequestMultiError, or nil if none found.
func (m *DeleteTrackRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTrackRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PlaylistId

	// no validation rules for TrackId

	if len(errors) > 0 {
		return DeleteTrackRequestMultiError(errors)
	}

	return nil
}

// DeleteTrackRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteTrackRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteTrackRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTrackRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTrackRequestMultiError) AllErrors() []error { return m }

// DeleteTrackRequestValidationError is the validation error returned by
// DeleteTrackRequest.Validate if the designated constraints aren't met.
type DeleteTrackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTrackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTrackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTrackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTrackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTrackRequestValidationError) ErrorName() string {
	return "DeleteTrackRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTrackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTrackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTrackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTrackRequestValidationError{}

// Validate checks the field values on DeleteTrackReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteTrackReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTrackReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTrackReplyMultiError, or nil if none found.
func (m *DeleteTrackReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTrackReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return DeleteTrackReplyMultiError(errors)
	}

	return nil
}

// DeleteTrackReplyMultiError is an error wrapping multiple validation errors
// returned by DeleteTrackReply.ValidateAll() if the designated constraints
// aren't met.
type DeleteTrackReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTrackReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTrackReplyMultiError) AllErrors() []error { return m }

// DeleteTrackReplyValidationError is the validation error returned by
// DeleteTrackReply.Validate if the designated constraints aren't met.
type DeleteTrackReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTrackReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTrackReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTrackReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTrackReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTrackReplyValidationError) ErrorName() string { return "DeleteTrackReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteTrackReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTrackReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTrackReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTrackReplyValidationError{}

// Validate checks the field values on GetTrackRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTrackRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTrackRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTrackRequestMultiError, or nil if none found.
func (m *GetTrackRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTrackRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PlaylistId

	// no validation rules for TrackId

	if len(errors) > 0 {
		return GetTrackRequestMultiError(errors)
	}

	return nil
}

// GetTrackRequestMultiError is an error wrapping multiple validation errors
// returned by GetTrackRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTrackRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTrackRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTrackRequestMultiError) AllErrors() []error { return m }

// GetTrackRequestValidationError is the validation error returned by
// GetTrackRequest.Validate if the designated constraints aren't met.
type GetTrackRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTrackRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTrackRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTrackRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTrackRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTrackRequestValidationError) ErrorName() string { return "GetTrackRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTrackRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTrackRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTrackRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTrackRequestValidationError{}

// Validate checks the field values on GetTrackReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTrackReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTrackReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTrackReplyMultiError, or
// nil if none found.
func (m *GetTrackReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTrackReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTrack()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTrackReplyValidationError{
					field:  "Track",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTrackReplyValidationError{
					field:  "Track",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTrack()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTrackReplyValidationError{
				field:  "Track",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetTrackReplyMultiError(errors)
	}

	return nil
}

// GetTrackReplyMultiError is an error wrapping multiple validation errors
// returned by GetTrackReply.ValidateAll() if the designated constraints
// aren't met.
type GetTrackReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTrackReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTrackReplyMultiError) AllErrors() []error { return m }

// GetTrackReplyValidationError is the validation error returned by
// GetTrackReply.Validate if the designated constraints aren't met.
type GetTrackReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTrackReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTrackReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTrackReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTrackReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTrackReplyValidationError) ErrorName() string { return "GetTrackReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetTrackReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTrackReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTrackReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTrackReplyValidationError{}

// Validate checks the field values on GetTracksRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTracksRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTracksRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTracksRequestMultiError, or nil if none found.
func (m *GetTracksRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTracksRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PlaylistId

	if len(errors) > 0 {
		return GetTracksRequestMultiError(errors)
	}

	return nil
}

// GetTracksRequestMultiError is an error wrapping multiple validation errors
// returned by GetTracksRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTracksRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTracksRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTracksRequestMultiError) AllErrors() []error { return m }

// GetTracksRequestValidationError is the validation error returned by
// GetTracksRequest.Validate if the designated constraints aren't met.
type GetTracksRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTracksRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTracksRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTracksRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTracksRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTracksRequestValidationError) ErrorName() string { return "GetTracksRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTracksRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTracksRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTracksRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTracksRequestValidationError{}

// Validate checks the field values on GetTracksReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetTracksReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTracksReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetTracksReplyMultiError,
// or nil if none found.
func (m *GetTracksReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTracksReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTracks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetTracksReplyValidationError{
						field:  fmt.Sprintf("Tracks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetTracksReplyValidationError{
						field:  fmt.Sprintf("Tracks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetTracksReplyValidationError{
					field:  fmt.Sprintf("Tracks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetTracksReplyMultiError(errors)
	}

	return nil
}

// GetTracksReplyMultiError is an error wrapping multiple validation errors
// returned by GetTracksReply.ValidateAll() if the designated constraints
// aren't met.
type GetTracksReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTracksReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTracksReplyMultiError) AllErrors() []error { return m }

// GetTracksReplyValidationError is the validation error returned by
// GetTracksReply.Validate if the designated constraints aren't met.
type GetTracksReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTracksReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTracksReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTracksReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTracksReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTracksReplyValidationError) ErrorName() string { return "GetTracksReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetTracksReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTracksReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTracksReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTracksReplyValidationError{}
