package filerepository

import (
	"FiLan/domain"
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
	// TODO: save to local file

	model := convertDomainToModel(file)
	return repo.DB.Create(&model).Error
}

// Delete is file deleting function from db
func (repo FileRepository) Delete(fullPath string) error {
	// TODO: delete from localfile
	path, fileName := convertFullPath(fullPath)

	result := repo.DB.Where("path = ?", path).Where("name = ?", fileName).Delete(&fileModel{})

	return result.Error
}

// GetByFullPath is file getter function from db
func (repo FileRepository) GetByFullPath(fullPath string) (domain.File, error) {
	// TODO: data get from localfile

	path, fileName := convertFullPath(fullPath)
	var receiver fileModel

	err := repo.DB.Where("path = ?", path).Where("name = ?", fileName).First(&receiver).Error
	if err != nil {
		return domain.File{}, err
	}

	data := make([]byte, receiver.Size)

	return convertModelToDomain(receiver, data), nil
}

// GetByDir is files getter function from db
func (repo FileRepository) GetByDir(path string) (files []domain.File, err error) {
	var receiver []fileModel
	err = repo.DB.Where("path = ?", path).Find(&receiver).Error
	if err != nil {
		return
	}

	for _, file := range receiver {
		// TODO: data get from localfile
		data := make([]byte, file.Size)
		files = append(files, convertModelToDomain(file, data))
	}

	return
}
