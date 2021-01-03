// Package filer provides file controll usecases
package filer

import (
	"time"
	"FiLan/domain"
	"FiLan/repository"
)

// Filer is struct implementing Filer usecase
type Filer struct {
	FileRepository repository.FileRepository
}

// New is constructor for Filer
func New(repository repository.FileRepository) *Filer {
	return &Filer{ FileRepository: repository }
}

// SaveFile is method saving file
func (filer *Filer) SaveFile(data []byte, name string, path string) (domain.File, error) {
	file := domain.File{
		Name: name,
		Path: path,
		Data: data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := filer.FileRepository.Save(file)

	return file, err
}

// DeleteFile is method deleteing file
func (filer *Filer) DeleteFile(name string, path string) (domain.File, error) {
	fullpath := joinPath(name, path)
	file, err := filer.FileRepository.GetByFullPath(fullpath)
	if err != nil {
		return file, err
	}
	err = filer.FileRepository.Delete(fullpath)

	return file, err
}

// GetFile is method getting file
func (filer *Filer) GetFile(name string, path string) (domain.File, error) {
	fullpath := joinPath(name, path)

	return filer.FileRepository.GetByFullPath(fullpath)
}

// GetFiles is method getting files by path
func (filer *Filer) GetFiles(path string) ([]domain.File, error) {
	return filer.FileRepository.GetByDir(path)
}

func joinPath(name string, path string) string {
	fullpath := path
	if (fullpath != "") {
		fullpath += "/"
	}
	fullpath += name

	return fullpath
}
