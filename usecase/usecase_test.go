package usecase_test

import (
	"testing"

	"FiLan/domain"
	"FiLan/usecase"
)

type testFiler struct{}

func (file testFiler) SaveFile(data []byte, name string, path string) domain.File {
	return domain.File{}
}

func (file testFiler) DeleteFile(name string, path string) domain.File {
	return domain.File{}
}

func (file testFiler) GetFile(name string, path string) domain.File {
	return domain.File{}
}

func assertFunction(saver usecase.FileSaver, deleter usecase.FileDeleter, getter usecase.FileGetter) {
}

func TestUsecaseDefinition(t *testing.T) {
	file := testFiler{}
	assertFunction(file, file, file)
}
