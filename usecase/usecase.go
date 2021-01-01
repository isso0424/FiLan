// Package usecase provides definition for usecase
package usecase

import "FiLan/domain"

// FileSaver is interface for file saving method
type FileSaver interface {
	SaveFile(data []byte, name string, path string) domain.File
}

// FileDeleter is interface for file deleting method
type FileDeleter interface {
	DeleteFile(name string, path string) domain.File
}

// FileGetter is interface for file getting
type FileGetter interface {
	Get(name string, path string) domain.File
}
