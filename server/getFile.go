package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func getFileHandler(w http.ResponseWriter, r *http.Request) {
	type Query struct {
		Name string
		Path string
	}

	const endpoint = "/file"
	const method = "GET"
	query := Query{}

	err := decoder.Decode(&query, r.URL.Query())
	if err != nil {
		queryNotEnoughError(w, endpoint, method, "name or path")

		return
	}

	file, err := controller.GetFile(query.Name, query.Path)
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
