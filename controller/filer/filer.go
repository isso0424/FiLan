// Package filer provides file control usecases
package filer

import (
	"FiLan/domain"
	"FiLan/repository"
	"FiLan/errors"
	"strings"
	"time"
)

// Filer is struct implementing Filer usecase
type Filer struct {
	FileRepository repository.FileRepository
}

// New is constructor for Filer
func New(repository repository.FileRepository) *Filer {
	return &Filer{FileRepository: repository}
}

func checkPathIsValid(path string) bool {
	splitted := strings.Split(path, "/")
	for _, value := range splitted {
		if value == "" {
			return false
		}
	}

	return true
}

// SaveFile is method saving file
func (filer *Filer) SaveFile(data []byte, name string, path string) (file domain.File, err error) {
	if strings.Contains(name, "/") || name == "" {
		err = errors.InvalidFileName{FileName: name}

		return
	}

	if len(data) == 0 {
		err = errors.EmptyData{}

		return
	}

	if !checkPathIsValid(path) {
		err = errors.InvalidFilePath{FilePath: path}

		return
	}

	file = domain.File{
		Name:      name,
		Path:      path,
		Data:      data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = filer.FileRepository.Save(file)

	return
}

// DeleteFile is method deleting file
func (filer *Filer) DeleteFile(name string, path string) (file domain.File, err error) {
	if strings.Contains(name, "/") || name == "" {
		err = errors.InvalidFileName{FileName: name}

		return
	}

	if !checkPathIsValid(path) {
		err = errors.InvalidFilePath{FilePath: path}

		return
	}

	fullpath := joinPath(name, path)
	file, err = filer.FileRepository.GetByFullPath(fullpath)
	if err != nil {
		return
	}
	err = filer.FileRepository.Delete(fullpath)

	return
}

// GetFile is method getting file
func (filer *Filer) GetFile(name string, path string) (file domain.File, err error) {
	if strings.Contains(name, "/") || name == "" {
		err = errors.InvalidFileName{FileName: name}

		return
	}

	if !checkPathIsValid(path) {
		err = errors.InvalidFilePath{FilePath: path}

		return
	}

	fullpath := joinPath(name, path)

	file, err = filer.FileRepository.GetByFullPath(fullpath)

	return
}

// GetFiles is method getting files by path
func (filer *Filer) GetFiles(path string) (files []domain.File, err error) {
	if !checkPathIsValid(path) {
		err = errors.InvalidFilePath{FilePath: path}

		return
	}

	files, err = filer.FileRepository.GetByDir(path)

	return
}

func joinPath(name string, path string) string {
	fullpath := path
	if fullpath != "" {
		fullpath += "/"
	}
	fullpath += name

	return fullpath
}
