package server

import (
	"FiLan/domain"
	"encoding/base64"
	"encoding/json"
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
		CreatedAt: file.CreatedAt,
		UpdatedAt: file.UpdatedAt,
	}
}

func handlerRequestError(w http.ResponseWriter, endpoint string, method string, statusCode int, errorMessage string) {
	log.Printf(logFormat, method, endpoint, statusCode, errorMessage)

	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(errorMessage))
	if err != nil {
		log.Println(err)
	}
}

func handleInternalServerError(w http.ResponseWriter, endpoint string, method string, occurredErr error) {
	log.Printf(logFormat, method, endpoint, http.StatusInternalServerError, occurredErr)

	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte("Internal server error"))
	if err != nil {
		log.Println(err)
	}
}

func domainWritebackToClient(file domain.File, w http.ResponseWriter, endpoint string, method string) {
	model := convertDomainToModel(file)
	encoded, err := json.Marshal(model)
	if err != nil {
		errorMessage := "error occur in json pasing"
		handlerRequestError(w, endpoint, method, http.StatusInternalServerError, errorMessage)

		return
	}

	_, err = w.Write(encoded)
	if err != nil {
		log.Println(err)
	}
}
