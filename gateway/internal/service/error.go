package service

import "fmt"

type errKind int

var (
	ErrCreatingRequestFailed = Error{kind: creatingRequestFailed}
	ErrRequestFailed         = Error{kind: requestFailed}
	ErrReadingBodyFailed     = Error{kind: readingBodyFailed}
	ErrParsingBodyFailed     = Error{kind: parsingBodyFailed}
	ErrUnknown               = Error{kind: unknown}
)

const (
	_ errKind = iota
	creatingRequestFailed
	requestFailed
	readingBodyFailed
	parsingBodyFailed
	unknown
)

type Error struct {
	kind errKind
	err  error
}

func (e *Error) Error() string {
	switch e.kind {
	case creatingRequestFailed:
		return fmt.Sprintf("Fail to create request %v", e.err)
	case requestFailed:
		return fmt.Sprintf("Request failed %v", e.err)
	case readingBodyFailed:
		return fmt.Sprintf("Reading body failed %v", e.err)
	case parsingBodyFailed:
		return fmt.Sprintf("Parsing body failed %v", e.err)
	default:
		return fmt.Sprintf("Unknown error %v", e.err)
	}
}

func (e *Error) Unwrap() error {
	return e.err
}

func (e *Error) Is(err error) bool {
	target, ok := err.(*Error)
	if !ok {
		return false
	}
	return target.kind == e.kind
}

func (e *Error) SetError(err error) *Error {
	e.err = err
	return e
}
