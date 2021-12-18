package mustache

import (
	"fmt"
)

const (
	ErrUnmatchedOpenTag      = "unmatched-open-tag"
	ErrEmptyTag              = "empty-tag"
	ErrSectionNoClosingTag   = "section-no-closing-tag"
	ErrInterleavedClosingTag = "interleaved-closing-tag"
	ErrInvalidMetaTag        = "invalid-meta-tag"
	ErrUnmatchedCloseTag     = "unmatched-close-tag"
)

type Error struct {
	Line    int
	Code    string
	Reason  string
}

func (e Error) Error() string {
	return e.defaultMessage()
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

func newError(line int, code string) Error {
	return Error{
		Line:    line,
		Code:    code,
		Reason:  "",
	}
}

func newErrorWithReason(line int, code, reason string) Error {
	return Error{
		Line:    line,
		Code:    code,
		Reason:  reason,
	}
}
