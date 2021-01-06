// Package repository provides definition for repository
package repository

import "FiLan/domain"

// FileRepository is file managing repository
type FileRepository interface {
	Save(file domain.File) error
	Delete(fullPath string) error
	GetByFullPath(path string) (domain.File, error)
	GetByDir(path string) ([]domain.File, error)
}

// FileAccessRepository is file "access" managing repository
type FileAccessRepository interface {
	Save(file domain.File) error
	Delete(fullPath string) error
	GetByFullPath(path string) ([]byte, error)
	GetByDir(path string) ([][]byte, error)
}
