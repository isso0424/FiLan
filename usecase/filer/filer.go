// Package filer provides file controll usecases
package filer

import (
	"FiLan/domain"
	"FiLan/repository"
)

// Filer is struct implementing Filer usecase
type Filer struct {
	FileRepository repository.FileRepository
}

// New is constructor for Filer
func New(repository repository.FileRepository) *Filer {
	return &Filer{ FileRepository: repository }
}

// SaveFile is method saving file
func (filer *Filer) SaveFile(data []byte, name string, path string) (domain.File, error) {
	return domain.File{}, nil
}

// DeleteFile is method deleteing file
func (filer *Filer) DeleteFile(name string, path string) (domain.File, error) {
	return domain.File{}, nil
}

// GetFile is method getting file
func (filer *Filer) GetFile(name string, path string) (domain.File, error) {
	return domain.File{}, nil
}

// GetFiles is method getting files by path
func (filer *Filer) GetFiles(path string) ([]domain.File, error) {
	return []domain.File{}, nil
}
