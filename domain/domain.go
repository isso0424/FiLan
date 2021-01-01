package domain

import "time"

// File is domain for file object
type File interface {
	getName() string
	getSize() uint64
	getPath() string

	getCreatedAt() time.Time
	getUpdatedAt() time.Time
}
