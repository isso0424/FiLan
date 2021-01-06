// Package accessrepository provides mock Repository for FileAccessRepository
package accessrepository

import (
	"FiLan/domain"
	"fmt"
)

// MockRepository is mock for FileAccessRepository
type MockRepository struct {
	files []domain.File
}

// New is constructor for MockRepository
func New() MockRepository {
	return MockRepository{files: []domain.File{}}
}

type invalidPathError struct {
	path string
}

func (err invalidPathError) Error() string {
	return fmt.Sprintf("The path %s is invalid", err.path)
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
