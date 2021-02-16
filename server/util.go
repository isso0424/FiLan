package server

import (
	"FiLan/model/domain"
)

func convertDomainToModel(file domain.File) fileModel {
	return fileModel{
		Name:      file.Name,
		Path:      file.Path,
		CreatedAt: file.CreatedAt,
		UpdatedAt: file.UpdatedAt,
	}
}
