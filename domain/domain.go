// Package domain is package for defining domains
package domain

import "time"

// File is domain struct for file object
type File struct {
	Name string
	Path string
	Data []byte

	CreatedAt time.Time
	UpdatedAt time.Time
}
