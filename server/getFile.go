package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func getFileHandler(w http.ResponseWriter, r *http.Request) {
	const endpoint = "/file"
	var name string
	var path string
	err := decoder.Decode(name, r.URL.Query())
	if err != nil {
		errorMessage := "Need query parameter: name"
		handlerRequestError(w, endpoint, http.StatusBadRequest, errorMessage)

		return
	}
	err = decoder.Decode(path, r.URL.Query())
	if err != nil {
		errorMessage := "Need query parameter: path"
		handlerRequestError(w, endpoint, http.StatusBadRequest, errorMessage)

		return
	}

	file, err := controller.GetFile(name, path)
	if err != nil {
		errorMessage := "Not found"
		handlerRequestError(w, endpoint, http.StatusNotFound, errorMessage)

		return
	}

	model := convertDomainToModel(file)
	encoded, err := json.Marshal(model)
	if err != nil {
		errorMessage := "error occur in json pasing"
		handlerRequestError(w, endpoint, http.StatusInternalServerError, errorMessage)

		return
	}

	_, err = w.Write(encoded)
	if err != nil {
		log.Println(err)
	}
}
