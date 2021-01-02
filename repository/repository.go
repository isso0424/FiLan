// Package repository provides definition for repository
package repository

import "FiLan/domain"

// FileRepository is file managing repository
type FileRepository interface {
	Save(file domain.File)
	Delete(file domain.File)
	GetByFullPath(path string) domain.File
	GetByDir(path string) []domain.File
}