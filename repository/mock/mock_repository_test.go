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
		Name:      "hoge",
		Path:      "foo/bar",
		Data:      []byte("example"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Save(file)
	if err != nil {
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
	if err == nil {
		t.Fatal("error should occur in file name being empty")
	}

	file.Name = "hoge"
	file.Path = ""
	err = repo.Save(file)
	if err != nil {
		t.Fatal(err)
	}

	file.Data = []byte{}
	err = repo.Save(file)
	if err == nil {
		t.Fatal("error should occur in data being empty")
	}
}

func TestSearchByFullPath(t *testing.T) {
	repo := mock.New()

	file := domain.File{
		Name:      "hoge",
		Path:      "foo/bar",
		Data:      []byte("example"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Save(file)
	if err != nil {
		t.Fatal(err)
	}

	searched, err := repo.GetByFullPath("foo/bar/hoge")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, searched.Name, repo.Files[0].Name)
	assert.Equal(t, searched.Path, repo.Files[0].Path)
	assert.Equal(t, string(searched.Data), string(repo.Files[0].Data))
	assert.Equal(t, searched.CreatedAt, repo.Files[0].CreatedAt)
	assert.Equal(t, searched.UpdatedAt, repo.Files[0].UpdatedAt)

	searched, err = repo.GetByFullPath("invalid/file")
	if err == nil {
		t.Fatal("error should occur with invalid path")
	}
}

func TestSearchByDir(t *testing.T) {
	repo := mock.New()

	file := domain.File{
		Name:      "hoge",
		Path:      "foo/bar",
		Data:      []byte("example"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Save(file)
	if err != nil {
		t.Fatal(err)
	}

	secondFile := domain.File{
		Name:      "fuga",
		Path:      "foo/bar",
		Data:      []byte("example"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = repo.Save(secondFile)
	if err != nil {
		t.Fatal(err)
	}

	searchedList, err := repo.GetByDir("foo/bar")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, searchedList[1].Name, repo.Files[1].Name)
	assert.Equal(t, searchedList[1].Path, repo.Files[1].Path)
	assert.Equal(t, string(searchedList[1].Data), string(repo.Files[1].Data))
	assert.Equal(t, searchedList[1].CreatedAt, repo.Files[1].CreatedAt)
	assert.Equal(t, searchedList[1].UpdatedAt, repo.Files[1].UpdatedAt)

	thirdFile := domain.File{
		Name:      "fuga",
		Path:      "bar/foo",
		Data:      []byte("example"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = repo.Save(thirdFile)
	if err != nil {
		t.Fatal(err)
	}

	searchedList, err = repo.GetByDir("foo/bar")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(searchedList), 2)
}

func TestDelete(t *testing.T) {
	repo := mock.New()

	file := domain.File{
		Name:      "hoge",
		Path:      "foo/bar",
		Data:      []byte("example"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.Save(file)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Delete("foo/bar/hoge")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(repo.Files), 0)
	err = repo.Delete("invalid/file")
	if err == nil {
		t.Fatal("error should occur with invalid target")
	}
}
