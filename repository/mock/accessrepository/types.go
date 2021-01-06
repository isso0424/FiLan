// Package accessrepository provides mock Repository for FileAccessRepository
package accessrepository

import (
	"FiLan/domain"
	"fmt"
)

// MockRepository is mock for FileAccessRepository
type MockRepository struct {
	Files []domain.File
}

// New is constructor for MockRepository
func New() MockRepository {
	return MockRepository{Files: []domain.File{}}
}

type notFoundError struct {
	name string
	dir  string
}

func (err notFoundError) Error() string {
	return fmt.Sprintf("The file %s is not found in %s", err.name, err.dir)
}

type notFoundErrorWithDir struct {
	dir string
}

func (err notFoundErrorWithDir) Error() string {
	return fmt.Sprintf("Any files does not exist in %s", err.dir)
}
