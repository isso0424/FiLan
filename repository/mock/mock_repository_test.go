package mock_test

import (
	"testing"
	"time"

	"FiLan/domain"
	"FiLan/repository/mock"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	repo := mock.New()

	assert.Equal(t, len(repo.Files), 0)
}

func TestSave(t *testing.T) {
	repo := mock.New()

	file := domain.File{
		Name: "hoge",
		Path: "foo/bar",
		Data: []byte("example"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Save(file)
	if (err != nil) {
		t.Fatal(err)
	}

	assert.Equal(t, len(repo.Files), 1)
	assert.Equal(t, file.Name, repo.Files[0].Name)
	assert.Equal(t, file.Path, repo.Files[0].Path)
	assert.Equal(t, string(file.Data), string(repo.Files[0].Data))
	assert.Equal(t, file.CreatedAt, repo.Files[0].CreatedAt)
	assert.Equal(t, file.UpdatedAt, repo.Files[0].UpdatedAt)

	file.Name = ""

	err = repo.Save(file)
	if (err == nil) {
		t.Fatal("error should occur in file name being empty")
	}

	file.Name = "hoge"
	file.Path = ""
	err = repo.Save(file)
	if (err != nil) {
		t.Fatal(err)
	}

	file.Data = []byte{}
	err = repo.Save(file)
	if (err == nil) {
		t.Fatal("error should occur in data being empty")
	}
}

func TestSearch(t *testing.T) {
	repo := mock.New()

	file := domain.File{
		Name: "hoge",
		Path: "foo/bar",
		Data: []byte("example"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Save(file)
	if (err != nil) {
		t.Fatal(err)
	}

	searched, err := repo.GetByFullPath("foo/bar/hoge")
	if (err != nil) {
		t.Fatal(err)
	}
	assert.Equal(t, searched.Name, repo.Files[0].Name)
	assert.Equal(t, searched.Path, repo.Files[0].Path)
	assert.Equal(t, string(searched.Data), string(repo.Files[0].Data))
	assert.Equal(t, searched.CreatedAt, repo.Files[0].CreatedAt)
	assert.Equal(t, searched.UpdatedAt, repo.Files[0].UpdatedAt)

	searchedList, err := repo.GetByDir("foo/bar")
	assert.Equal(t, searchedList[0].Name, repo.Files[0].Name)
	assert.Equal(t, searchedList[0].Path, repo.Files[0].Path)
	assert.Equal(t, string(searchedList[0].Data), string(repo.Files[0].Data))
	assert.Equal(t, searchedList[0].CreatedAt, repo.Files[0].CreatedAt)
	assert.Equal(t, searchedList[0].UpdatedAt, repo.Files[0].UpdatedAt)
}

func TestDelete(t *testing.T) {
	repo := mock.New()

	file := domain.File{
		Name: "hoge",
		Path: "foo/bar",
		Data: []byte("example"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Save(file)
	if (err != nil) {
		t.Fatal(err)
	}

	err = repo.Delete("foo/bar/hoge")
	if (err != nil) {
		t.Fatal(err)
	}

	assert.Equal(t, len(repo.Files), 0)
}
