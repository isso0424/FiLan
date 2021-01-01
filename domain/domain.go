// Package domain is package for defining domains
package domain

import "time"

// File is domain for file object
type File interface {
	GetName() string
	GetSize() uint64
	GetPath() string

	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
}
