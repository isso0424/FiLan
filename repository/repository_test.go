package repository_test

import (
	"FiLan/domain"
	"FiLan/repository"
	"testing"
)

type testRepository struct {
}

func (repository testRepository) Save(file domain.File) {
}

func (repository testRepository) Delete(file domain.File) {
}

func (repository testRepository) GetByFullPath(path string) domain.File {
	return domain.File{}
}

func (repository testRepository) GetByDir(path string) []domain.File {
	return []domain.File{}
}

func assertFunction(repository repository.FileRepository) {}

func TestRepository(t *testing.T) {
	repository := testRepository{}
	assertFunction(repository)
}
