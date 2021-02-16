package filerepository

import (
	"FiLan/model/domain"
	"strings"
)

func convertDomainToModel(file domain.File) fileModel {
	return fileModel{
		Name:        file.Name,
		Size:        len(file.Data),
		Path:        file.Path,
		AddedAt:     file.CreatedAt,
		RefreshedAt: file.UpdatedAt,
	}
}

func convertModelToDomain(model fileModel) domain.File {
	return domain.File{
		Name:      model.Name,
		Path:      model.Path,
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

	return convertModelToDomain(receiver), nil
}

// GetByDir is files getter function from db
func (repo FileRepository) GetByDir(dir string) (files []domain.File, err error) {
	var receiver []fileModel
	err = repo.DB.Where("path = ?", dir).Find(&receiver).Error
	if err != nil {
		return
	}

	for _, file := range receiver {
		files = append(files, convertModelToDomain(file))
	}

	return
}
