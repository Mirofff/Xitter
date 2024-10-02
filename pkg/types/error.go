package types

import (
	"errors"
)

// IllegalArgumentException - common error for any function argument processing.
type IllegalArgumentException = error

// TODOException - common error yet not defined
type TODOException = error

// NewIllegalArgumentException - return new IllegalArgumentException
func NewIllegalArgumentException(msg string) IllegalArgumentException {
	return errors.New(msg)
}

// NewTODOException - return new TODOException
func NewTODOException(msg string) TODOException {
	return errors.New(msg)
}
