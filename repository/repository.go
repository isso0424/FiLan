package repository

import "FiLan/domain"

// FileRepository is file managing repository
type FileRepository interface {
	Save(file domain.File)
	Delete(file domain.File)
	GetByName(path string) domain.File
	GetByDir(path string) []domain.File
}
