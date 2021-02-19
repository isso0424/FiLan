package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func getFilesHandler(w http.ResponseWriter, r *http.Request) {
	type Query struct {
		Path string `schema:"path,required"`
	}

	const endpoint = "/files"
	const method = "GET"
	query := Query{}

	err := decoder.Decode(&query, r.URL.Query())
	if err != nil {
		handleInvalidQuery(w, endpoint, method, "path")

		return
	}

	files, err := controller.GetFiles(query.Path)
	if err != nil {
		handleInternalServerError(w, endpoint, method, err)

		return
	}

	filesModel := []fileModel{}
	for _, file := range files {
		fileModel := convertDomainToModel(file)
		filesModel = append(filesModel, fileModel)
	}

	encoded, err := json.Marshal(filesModel)
	if err != nil {
		handleJSONParseError(w, endpoint, method)

		return
	}

	_, err = w.Write(encoded)
	if err != nil {
		log.Println(err)
	}

	loggingSuccess(method, endpoint, http.StatusOK)
}
