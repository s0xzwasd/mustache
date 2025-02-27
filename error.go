package mustache

import (
	"fmt"
)

// ErrorCode is the list of allowed values for the error's code.
type ErrorCode string

// List of values that ErrorCode can take.
const (
	ErrUnmatchedOpenTag      ErrorCode = "unmatched_open_tag"
	ErrEmptyTag              ErrorCode = "empty_tag"
	ErrSectionNoClosingTag   ErrorCode = "section_no_closing_tag"
	ErrInterleavedClosingTag ErrorCode = "interleaved_closing_tag"
	ErrInvalidMetaTag        ErrorCode = "invalid_meta_tag"
	ErrUnmatchedCloseTag     ErrorCode = "unmatched_close_tag"
)

// Error represents an error during the parsing
type Error struct {
	// Line returns the line of the error
	Line int
	// Code returns the error code of the error
	Code ErrorCode
	// returns the name of the element generating the error
	Reason string
}

func (e Error) Error() string {
	return fmt.Sprintf("line %d: %s", e.Line, e.defaultMessage())
}

func (e Error) defaultMessage() string {
	switch e.Code {
	case ErrUnmatchedOpenTag:
		return "unmatched open tag"
	case ErrEmptyTag:
		return "empty tag"
	case ErrSectionNoClosingTag:
		return fmt.Sprintf("Section %s has no closing tag", e.Reason)
	case ErrInterleavedClosingTag:
		return fmt.Sprintf("interleaved closing tag: %s", e.Reason)
	case ErrInvalidMetaTag:
		return "Invalid meta tag"
	case ErrUnmatchedCloseTag:
		return "unmatched close tag"
	default:
		return "unknown error"
	}
}

func newError(line int, code ErrorCode) Error {
	return Error{
		Line:   line,
		Code:   code,
		Reason: "",
	}
}

func newErrorWithReason(line int, code ErrorCode, reason string) Error {
	return Error{
		Line:   line,
		Code:   code,
		Reason: reason,
	}
}
