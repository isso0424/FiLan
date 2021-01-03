package server

import (
	"FiLan/domain"
	"encoding/base64"
	"log"
	"net/http"
)

// nolint:unused,deadcode
func convertModelToDomain(model fileModel) domain.File {
	decoded, _ := base64.RawStdEncoding.DecodeString(model.Encoded)

	return domain.File{
		Name:      model.Name,
		Path:      model.Path,
		Data:      decoded,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func convertDomainToModel(file domain.File) fileModel {
	return fileModel{
		Name:      file.Name,
		Path:      file.Path,
		Encoded:   base64.StdEncoding.EncodeToString(file.Data),
		CreatedAt: file.CreatedAt,
		UpdatedAt: file.UpdatedAt,
	}
}

// nolint:unparam
func handlerRequestError(w http.ResponseWriter, endpoint string, statusCode int, errorMessage string) {
	log.Printf(logFormat, endpoint, http.StatusBadRequest, errorMessage)

	w.WriteHeader(http.StatusBadRequest)
	_, err := w.Write([]byte(errorMessage))
	if err != nil {
		log.Println(err)
	}
}
