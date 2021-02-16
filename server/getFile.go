package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func getFileHandler(w http.ResponseWriter, r *http.Request) {
	const endpoint = "/file"
	const method = "GET"
	var name string
	var path string
	err := decoder.Decode(name, r.URL.Query())
	if err != nil {
		queryNotEnoughError(w, endpoint, method, "name")

		return
	}

	err = decoder.Decode(path, r.URL.Query())
	if err != nil {
		queryNotEnoughError(w, endpoint, method, "path")

		return
	}

	file, err := controller.GetFile(name, path)
	if err != nil {
		errorMessage := "Not found"
		handlerRequestError(w, endpoint, method, http.StatusNotFound, errorMessage)

		return
	}

	model := convertDomainToModel(file)
	encoded, err := json.Marshal(model)
	if err != nil {
		jsonParseError(w, endpoint, method)

		return
	}

	_, err = w.Write(encoded)
	if err != nil {
		log.Println(err)
	}
}
