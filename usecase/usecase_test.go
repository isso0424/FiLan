package usecase_test

import (
	"FiLan/domain"
	"FiLan/usecase"
	"testing"
	"time"
)

type testFiler struct {
}

type testFile struct {
	Name      string
	Size      uint64
	Path      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (file testFile) GetName() string {
	return file.Name
}

func (file testFile) GetSize() uint64 {
	return file.Size
}

func (file testFile) GetPath() string {
	return file.Path
}

func (file testFile) GetCreatedAt() time.Time {
	return file.CreatedAt
}

func (file testFile) GetUpdatedAt() time.Time {
	return file.UpdatedAt
}

func (file testFiler) SaveFile(data []byte, name string, path string) domain.File {
	return testFile{}
}

func (file testFiler) DeleteFile(name string, path string) domain.File {
	return testFile{}
}

func (file testFiler) GetFile(name string, path string) domain.File {
	return testFile{}
}

func assertFunction(saver usecase.FileSaver, deleter usecase.FileDeleter, getter usecase.FileGetter) {
}

func TestUsecaseDefinition(t *testing.T) {
	file := testFiler{}
	assertFunction(file, file, file)
}
