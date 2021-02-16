package filer_test

import (
	"FiLan/controller/filer"
	"FiLan/repository/mock/db"
	"FiLan/repository/mock/filesystem"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	name = "hoge"
	path = "foo/bar"
)

type testError struct {
	Message string
}

func (err testError) Error() string {
	return err.Message
}

func TestSavingFileSuccess(t *testing.T) {
	accessRepository := filesystem.New()
	mockRepository := db.New()
	controller := filer.New(mockRepository, &accessRepository)
	data := []byte("example")

	file, err := controller.SaveFile(data, name, path)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, name, file.Name)
	assert.Equal(t, path, file.Path)
	assert.Equal(t, string(data), string(file.Data))
}

func TestSavingFileFail(t *testing.T) {
	accessRepository := filesystem.New()
	mockRepository := db.New()
	controller := filer.New(mockRepository, &accessRepository)
	data := []byte("example")

	invalidData := []byte{}
	_, err := controller.SaveFile(invalidData, name, path)
	if err == nil {
		t.Fatal(testError{Message: "error should occur with invalid data length: 0"})
	}

	invalidName := ""
	_, err = controller.SaveFile(data, invalidName, path)
	if err == nil {
		t.Fatal(testError{Message: "error should occur with invalid name: empty string"})
	}
}

func TestDeletingSuccess(t *testing.T) {
	accessRepository := filesystem.New()
	mockRepository := db.New()
	controller := filer.New(mockRepository, &accessRepository)
	data := []byte("example")

	_, err := controller.SaveFile(data, name, path)
	if err != nil {
		t.Fatal(err)
	}

	_, err = controller.DeleteFile(name, path)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeletingFail(t *testing.T) {
	accessRepository := filesystem.New()
	mockRepository := db.New()
	controller := filer.New(mockRepository, &accessRepository)
	invalidName := "invalid"
	invalidPath := "example/invalid"
	_, err := controller.DeleteFile(invalidName, invalidPath)
	if err == nil {
		t.Fatal(testError{Message: "error should occur with invalid full path"})
	}
}

func TestGettingSuccess(t *testing.T) {
	accessRepository := filesystem.New()
	mockRepository := db.New()
	controller := filer.New(mockRepository, &accessRepository)
	data := []byte("example")

	_, err := controller.SaveFile(data, name, path)
	if err != nil {
		t.Fatal(err)
	}

	file, err := controller.GetFile(name, path)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(data), string(file.Data))
	assert.Equal(t, name, file.Name)
	assert.Equal(t, path, file.Path)
}

func TestGettingFail(t *testing.T) {
	accessRepository := filesystem.New()
	mockRepository := db.New()
	controller := filer.New(mockRepository, &accessRepository)
	invalidName := "invalid"
	invalidPath := "invalid/path"

	_, err := controller.GetFile(invalidName, invalidPath)
	if err == nil {
		t.Fatal(testError{Message: "error occur with invalid path"})
	}
}

func TestGettingFilesSuccess(t *testing.T) {
	accessRepository := filesystem.New()
	mockRepository := db.New()
	controller := filer.New(mockRepository, &accessRepository)
	data := []byte("example")

	_, err := controller.SaveFile(data, name, path)
	if err != nil {
		t.Fatal(err)
	}

	secondData := []byte("example2")
	secondName := "fuga"
	_, err = controller.SaveFile(secondData, secondName, path)
	if err != nil {
		t.Fatal(err)
	}

	thirdData := []byte("example3")
	thirdName := "foo"
	thirdPath := "bar"
	_, err = controller.SaveFile(thirdData, thirdName, thirdPath)
	if err != nil {
		t.Fatal(err)
	}

	files, err := controller.GetFiles(path)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(files), 2)

	assert.Equal(t, name, files[0].Name)
	assert.Equal(t, path, files[0].Path)
	assert.Equal(t, string(data), string(files[0].Data))

	assert.Equal(t, secondName, files[1].Name)
	assert.Equal(t, path, files[1].Path)
	assert.Equal(t, string(secondData), string(files[1].Data))
}
