package repository

import "fmt"

type errKind int

var (
	ErrDataNotFound = Error{kind: dataNotFound}
	ErrUnknown      = Error{kind: unknown}
)

const (
	_ errKind = iota
	dataNotFound
	unknown
)

type Error struct {
	kind errKind
	err  error
}

func (e *Error) Error() string {
	switch e.kind {
	case dataNotFound:
		return "Data not found"
	default:
		return fmt.Sprintf("Unknown Error %v", e.err)
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
