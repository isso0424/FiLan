// Package db provides mock struct for FileRepository
package db

import (
	"FiLan/model/domain"
	"strings"
)

// FileRepository is mock struct for FileRepository
type FileRepository struct {
	Files []domain.File
}

// New is constructor for Mock FileRepository
func New() *FileRepository {
	return &FileRepository{Files: []domain.File{}}
}

type err struct {
	Message string
}

func (e err) Error() string {
	return e.Message
}

func convertFullPath(fullpath string) (string, string) {
	dirs := strings.Split(fullpath, "/")
	fileName := dirs[len(dirs)-1]
	path := ""
	for index, dir := range dirs {
		if index == len(dirs)-1 {
			break
		}
		if index != 0 {
			path += "/"
		}
		path += dir
	}

	return path, fileName
}

// Save is file save function
func (repo *FileRepository) Save(file domain.File) error {
	if len(file.Data) == 0 {
		return err{Message: "file size 0 is not supported"}
	}
	if file.Name == "" {
		return err{Message: "file must not empty string"}
	}
	repo.Files = append(repo.Files, file)

	return nil
}

// Delete is file deleting function
func (repo *FileRepository) Delete(fullPath string) error {
	path, fileName := convertFullPath(fullPath)
	for index, file := range repo.Files {
		if file.Path == path && file.Name == fileName {
			repo.Files[index] = repo.Files[len(repo.Files)-1]
			repo.Files = repo.Files[:len(repo.Files)-1]

			return nil
		}
	}

	return err{Message: "target does not found"}
}

// GetByFullPath is file getting function by fullpath
func (repo *FileRepository) GetByFullPath(fullPath string) (domain.File, error) {
	path, fileName := convertFullPath(fullPath)
	for _, file := range repo.Files {
		if file.Path == path && file.Name == fileName {
			return file, nil
		}
	}

	return domain.File{}, err{Message: "target does not found"}
}

// GetByDir is file getting function by dir
func (repo *FileRepository) GetByDir(path string) ([]domain.File, error) {
	var result []domain.File
	for _, file := range repo.Files {
		if file.Path == path {
			result = append(result, file)
		}
	}

	return result, nil
}
