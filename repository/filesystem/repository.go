package filesystem

import (
	"FiLan/domain"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

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

// Save is method saving file to local storage
func (repo Repository) Save(file domain.File) error {
	filePath := path.Join(repo.StorageDir, file.Path, file.Name)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	_, err = f.Write(file.Data)

	return err
}

// Delete is method deleting file from local storage
func (repo Repository) Delete(file domain.File) error {
	filePath := path.Join(repo.StorageDir, file.Path, file.Name)
	return os.Remove(filePath)
}

// GetByFullPath is method getting file from local storage
func (repo Repository) GetByFullPath(fullpath string) (domain.File, error) {
	fullPath := path.Join(repo.StorageDir, fullpath)
	f, err := os.Open(fullPath)
	if err != nil {
		return domain.File{}, err
	}

	dir, name := convertFullPath(fullpath)

	file := domain.File{
		Name: name,
		Path: dir,
	}

	_, err = f.Read(file.Data)
	if err != nil {
		return file, err
	}

	return file, nil
}

// GetByDir is method getting files from local storage
func (repo Repository) GetByDir(dir string) (domains []domain.File, err error) {
	fullPath := path.Join(repo.StorageDir, dir)
	files, err := ioutil.ReadDir(fullPath)
	if err != nil {
		return
	}

	for _, file := range files {
		newFile := domain.File{
			Name: file.Name(),
			Path: dir,
		}
		filePath := path.Join(fullPath, file.Name())
		f, err := os.Open(filePath)
		if err != nil {
			return domains, err
		}
		_, err = f.Read(newFile.Data)
		if err != nil {
			return domains, err
		}

		domains = append(domains, newFile)
	}

	return
}
