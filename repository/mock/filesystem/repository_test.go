package filesystem_test

import (
	"FiLan/model/domain"
	"FiLan/repository/mock/filesystem"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestingFiles(repo *filesystem.MockRepository) (err error) {
	file := domain.File{
		Name: "hoge",
		Path: "foo/bar",
		Data: []byte("fuga"),
	}
	err = repo.Save(file)
	if err != nil {
		return
	}

	file2 := domain.File{
		Name: "fuga",
		Path: "foo/bar",
		Data: []byte("hoge"),
	}
	err = repo.Save(file2)
	if err != nil {
		return
	}

	return
}

func TestFileSavingSuccess(t *testing.T) {
	repo := filesystem.New()
	err := createTestingFiles(&repo)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hoge", repo.Files[0].Name)
	assert.Equal(t, "foo/bar", repo.Files[0].Path)
	assert.Equal(t, "fuga", string(repo.Files[0].Data))
	assert.Equal(t, "fuga", repo.Files[1].Name)
	assert.Equal(t, "foo/bar", repo.Files[1].Path)
	assert.Equal(t, "hoge", string(repo.Files[1].Data))

	file3 := domain.File{
		Name: "hoge",
		Path: "foo/bar",
		Data: []byte("hoge"),
	}

	err = repo.Save(file3)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hoge", repo.Files[0].Name)
	assert.Equal(t, "foo/bar", repo.Files[0].Path)
	assert.Equal(t, "hoge", string(repo.Files[0].Data))
}

func TestFileDeletingSuccess(t *testing.T) {
	repo := filesystem.New()
	err := createTestingFiles(&repo)
	if err != nil {
		t.Fatal(err)
	}
	err = repo.Delete("foo/bar/hoge")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(repo.Files))
	assert.Equal(t, "fuga", repo.Files[0].Name)
	assert.Equal(t, "hoge", string(repo.Files[0].Data))

	err = repo.Delete("foo/bar/fuga")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 0, len(repo.Files))
}

func TestFileGettingByFullPathSuccess(t *testing.T) {
	repo := filesystem.New()
	err := createTestingFiles(&repo)
	if err != nil {
		t.Fatal(err)
	}
	data, err := repo.GetByFullPath("foo/bar/hoge")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(repo.Files))
	assert.Equal(t, "fuga", string(data))

	file3 := domain.File{
		Name: "fuga",
		Path: "bar",
		Data: []byte("hogehoge"),
	}
	err = repo.Save(file3)
	if err != nil {
		t.Fatal(err)
	}

	data, err = repo.GetByFullPath("bar/fuga")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "hogehoge", string(data))
}

func TestFileGettingByDirSuccess(t *testing.T) {
	repo := filesystem.New()
	err := createTestingFiles(&repo)
	if err != nil {
		t.Fatal(err)
	}
	file3 := domain.File{
		Name: "fuga",
		Path: "bar/foo",
		Data: []byte("hoge"),
	}

	err = repo.Save(file3)
	if err != nil {
		t.Fatal(err)
	}

	dates, err := repo.GetByDir("foo/bar")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 2, len(dates))
	assert.Equal(t, "fuga", string(dates[0]))
	assert.Equal(t, "hoge", string(dates[1]))
}

func TestFileDeletingFail(t *testing.T) {
	repo := filesystem.New()
	err := createTestingFiles(&repo)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Delete("invalid")
	if err == nil {
		t.Fatal("Error should occur with invalid path")
	}
}

func TestFileGettingByFullPathFail(t *testing.T) {
	repo := filesystem.New()
	err := createTestingFiles(&repo)
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.GetByFullPath("invalid")
	if err == nil {
		t.Fatal("Error should occur with invalid path")
	}
}

func TestFileGettingByDirFail(t *testing.T) {
	repo := filesystem.New()
	err := createTestingFiles(&repo)
	if err != nil {
		t.Fatal(err)
	}

	_, err = repo.GetByDir("invalid")
	if err == nil {
		t.Fatal("Error should occur with invalid path")
	}
}
