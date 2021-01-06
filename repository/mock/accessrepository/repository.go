package accessrepository

import (
	"FiLan/domain"
	"fmt"
	"strings"
)

// Save is mock function for FileAccessRepository.Save
func (repo *MockRepository) Save(file domain.File) error {
	replaced := false
	for index, f := range repo.Files {
		if (file.Name == f.Name && file.Path == f.Path) {
			replaced = true
			repo.Files[index] = file
			break
		}
	}

	if (!replaced) {
		repo.Files = append(repo.Files, file)
	}

	return nil
}

// Delete is mock function for FileAccessRepository.Delete
func (repo *MockRepository) Delete(fullPath string) error {
	dir, name := convertFullPath(fullPath)
	finalIndex := len(repo.Files) - 1

	for index, f := range repo.Files {
		if f.Path == dir && f.Name == name {
			repo.Files[index] = repo.Files[finalIndex]
			repo.Files = repo.Files[:finalIndex]

			return nil
		}
	}

	return notFoundError{name: name, dir: dir}
}

// GetByFullPath is mock function for FileAccessRepository.GetByFullPath
func (repo MockRepository) GetByFullPath(fullPath string) ([]byte, error) {
	dir, name := convertFullPath(fullPath)

	for _, file := range repo.Files {
		if file.Path == dir && file.Name == name {
			return file.Data, nil
		}
	}

	return nil, notFoundError{name: name, dir: dir}
}

// GetByDir is mock function for FileAccessRepository.GetByDir
func (repo MockRepository) GetByDir(dir string) (dates [][]byte, err error) {
	for _, file := range repo.Files {
		if file.Path == dir {
			dates = append(dates, file.Data)
		}
	}

	if len(dates) == 0 {
		err = notFoundErrorWithDir{dir: dir}
	}

	return
}

func convertFullPath(fullPath string) (path string, name string) {
	dirs := strings.Split(fullPath, "/")
	finalIndex := len(dirs) - 1
	switch len(dirs) {
	case 0:
		break
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
