package testutil

import "errors"

var (
	ErrZeroValue = errors.New("zero value")
	ErrTooLong   = errors.New("too long")
)
