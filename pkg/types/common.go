// Package types - a commonly used package in the project, where popular types are defined.
package types

import (
	"github.com/google/uuid"
)

// GUID - identifying globally unique identifiers.
type GUID = uuid.UUID

// NewGUID - return new generated GUID.
func NewGUID() GUID {
	return uuid.New()
}

// SequenceID - sequence unsigned identifiers.
type SequenceID = uint64
