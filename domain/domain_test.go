package domain_test

import (
	"testing"
	"time"

	"FiLan/domain"
)

type testFile struct {
	Name string
	Size uint64
	Path string
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

func assertFunction(file domain.File) {
	return
}

func TestFileDomain(t *testing.T) {
	file := testFile{
		Name: "example",
		Size: 100,
		Path: "foo/bar",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	assertFunction(file)
}
