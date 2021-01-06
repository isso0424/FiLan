package accessrepository

import (
	"FiLan/domain"
	"fmt"
	"strings"
)

// Save is mock function for FileAccessRepository.Save
func (repo *MockRepository) Save(file domain.File) error {
	replaced := false
	for index, f := range repo.files {
		if (file.Name == f.Name && file.Path == f.Path) {
			replaced = true
			repo.files[index] = file
			break
		}
	}

	if (!replaced) {
		repo.files = append(repo.files, file)
	}

	return nil
}

// Delete is mock function for FileAccessRepository.Delete
func (repo *MockRepository) Delete(fullPath string) error {
	dir, name, err := convertFullPath(fullPath)
	if err != nil {
		return err
	}

	finalIndex := len(repo.files) - 1

	for index, f := range repo.files {
		if f.Path == dir && f.Name == name {
			repo.files[index] = repo.files[finalIndex]

			return nil
		}
	}

	return notFoundError{name: name, dir: dir}
}

// GetByFullPath is mock function for FileAccessRepository.GetByFullPath
func (repo MockRepository) GetByFullPath(fullPath string) ([]byte, error) {
	dir, name, err := convertFullPath(fullPath)
	if err != nil {
		return []byte{}, err
	}

	for _, file := range repo.files {
		if file.Path == dir && file.Name == name {
			return file.Data, nil
		}
	}

	return nil, notFoundError{name: name, dir: dir}
}

// GetByDir is mock function for FileAccessRepository.GetByDir
func (repo MockRepository) GetByDir(dir string) (dates [][]byte, err error) {
	for _, file := range repo.files {
		if file.Path == dir {
			dates = append(dates, file.Data)
		}
	}

	if len(dates) != 0 {
		err = notFoundErrorWithDir{dir: dir}
	}

	return
}

func convertFullPath(fullPath string) (path string, name string, err error) {
	dirs := strings.Split(fullPath, "/")
	finalIndex := len(dirs) - 1
	switch len(dirs) {
	case 0:
		err = invalidPathError{path: fullPath}
	case 1:
		name = dirs[0]
	default:
		for index, value := range dirs {
			switch index {
			case finalIndex:
				name = value
			case 0:
				path = value
			default:
				path += fmt.Sprintf("/%s", value)
			}
		}
	}

	return
}
