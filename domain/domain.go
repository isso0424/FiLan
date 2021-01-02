// Package domain is package for defining domains
package domain

import "time"

// File is domain struct for file object
type File struct {
	Name string
	Size uint64
	Path string

	CreatedAt time.Time
	UpdatedAt time.Time
}
