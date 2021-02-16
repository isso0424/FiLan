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

func domainWritebackToClient(file domain.File, w http.ResponseWriter, endpoint string, method string) {
	model := convertDomainToModel(file)
	encoded, err := json.Marshal(model)
	if err != nil {
		handleInternalServerError(w, endpoint, method, err)

		return
	}

	_, err = w.Write(encoded)
	if err != nil {
		log.Println(err)
	}
}
