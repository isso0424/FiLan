package server

import (
	"FiLan/model/domain"
	"encoding/json"
	"log"
	"net/http"
)

func writeError(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)

	_, err := w.Write([]byte(errorMessage))
	if err != nil {
		log.Println(err)
	}
}

func domainWritebackToClient(file domain.File, w http.ResponseWriter, endpoint string, method string) error {
	model := convertDomainToModel(file)
	encoded, err := json.Marshal(model)
	if err != nil {
		return err
	}

	_, err = w.Write(encoded)

	return err
}

func fileWritebackToClient(file []byte, w http.ResponseWriter, endpoint string, method string) error {
	_, err := w.Write(file)

	return err
}
