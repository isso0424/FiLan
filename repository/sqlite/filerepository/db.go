package filerepository

import (
	"FiLan/domain"
	"path"
	"strings"
)

func convertDomainToModel(file domain.File) fileModel {
	return fileModel{
		Name: file.Name,
		Size: len(file.Data),
		Path: file.Path,
		AddedAt: file.CreatedAt,
		RefreshedAt: file.UpdatedAt,
	}
}

func convertModelToDomain(model fileModel, data []byte) domain.File {
	return domain.File{
		Name: model.Name,
		Data: data,
		Path: model.Path,
		CreatedAt: model.AddedAt,
		UpdatedAt: model.RefreshedAt,
	}
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

// Save is file saving function to db
func (repo FileRepository) Save(file domain.File) error {
	err := repo.FileSystemRepository.Save(file)
	if err != nil {
		return err
	}

	model := convertDomainToModel(file)
	return repo.DB.Create(&model).Error
}

// Delete is file deleting function from db
func (repo FileRepository) Delete(fullPath string) error {
	path, fileName := convertFullPath(fullPath)

	err := repo.DB.Where("path = ?", path).Where("name = ?", fileName).Delete(&fileModel{}).Error
	if err != nil {
		return err
	}

	err = repo.FileSystemRepository.Delete(fullPath)

	return err
}

// GetByFullPath is file getter function from db
func (repo FileRepository) GetByFullPath(fullPath string) (domain.File, error) {
	path, fileName := convertFullPath(fullPath)
	var receiver fileModel

	err := repo.DB.Where("path = ?", path).Where("name = ?", fileName).First(&receiver).Error
	if err != nil {
		return domain.File{}, err
	}

	file, err := repo.FileSystemRepository.GetByFullPath(fullPath)
	if err != nil {
		return domain.File{}, err
	}

	return convertModelToDomain(receiver, file.Data), nil
}

// GetByDir is files getter function from db
func (repo FileRepository) GetByDir(dir string) (files []domain.File, err error) {
	var receiver []fileModel
	err = repo.DB.Where("path = ?", dir).Find(&receiver).Error
	if err != nil {
		return
	}

	for _, file := range receiver {
		filePath := path.Join(dir, file.Name)
		fileData, err := repo.FileSystemRepository.GetByFullPath(filePath)
		if err != nil {
			return files, err
		}
		files = append(files, convertModelToDomain(file, fileData.Data))
	}

	return
}
