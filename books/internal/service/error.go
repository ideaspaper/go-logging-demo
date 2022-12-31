package service

import "fmt"

type errKind int

var (
	ErrBookNotFound = Error{kind: bookNotFound}
	ErrUnknown      = Error{kind: unknown}
)

const (
	_ errKind = iota
	bookNotFound
	unknown
)

type Error struct {
	kind errKind
	err  error
}

func (e *Error) Error() string {
	switch e.kind {
	case bookNotFound:
		return fmt.Sprintf("Book not found %v", e.err)
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
