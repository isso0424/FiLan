// Package usecase provides definition for usecase
package usecase

import "FiLan/domain"

// Filer is interface for file managing
type Filer interface {
	SaveFile(data []byte, name string, path string) (domain.File, error)
	DeleteFile(name string, path string) (domain.File, error)
	GetFile(name string, path string) (domain.File, error)
}
