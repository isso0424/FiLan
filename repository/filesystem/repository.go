package filesystem

import (
	"FiLan/model/domain"
	"io/ioutil"
	"os"
	"path"
)

// Save is method saving file to local storage
func (repo Repository) Save(file domain.File) error {
	fileDir := path.Join(repo.StorageDir, file.Path)
	err := os.MkdirAll(fileDir, 0777)
	if err != nil {
		return err
	}

	filePath := path.Join(fileDir, file.Name)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}

	_, err = f.Write(file.Data)

	return err
}

// Delete is method deleting file from local storage
func (repo Repository) Delete(fullPath string) error {
	filePath := path.Join(repo.StorageDir, fullPath)

	return os.Remove(filePath)
}

// GetByFullPath is method getting file from local storage
func (repo Repository) GetByFullPath(fullpath string) (data []byte, err error) {
	fullPath := path.Join(repo.StorageDir, fullpath)
	data, err = ioutil.ReadFile(fullPath)

	return
}

// GetByDir is method getting files from local storage
func (repo Repository) GetByDir(dir string) (data [][]byte, err error) {
	fullPath := path.Join(repo.StorageDir, dir)
	files, err := ioutil.ReadDir(fullPath)
	if err != nil {
		return
	}

	for _, file := range files {
		newFile := []byte{}
		filePath := path.Join(fullPath, file.Name())
		f, err := os.Open(filePath)
		if err != nil {
			return data, err
		}
		_, err = f.Read(newFile)
		if err != nil {
			return data, err
		}

		data = append(data, newFile)
	}

	return
}
