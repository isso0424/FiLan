package server

import (
	"encoding/json"
	"fmt"
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
	queries := []loggingQuery{{key: "path", value: query.Path}}
	if err != nil {
		handleRequestError(w, endpoint, method, http.StatusBadRequest, queries, fmt.Sprintf(notEnoughQuery, "path"))

		return
	}

	files, err := controller.GetFiles(query.Path)
	if err != nil {
		handleRequestError(w, endpoint, method, http.StatusNotFound, queries, "Directory not found")

		return
	}

	filesModel := []fileModel{}
	for _, file := range files {
		fileModel := convertDomainToModel(file)
		filesModel = append(filesModel, fileModel)
	}

	encoded, err := json.Marshal(filesModel)
	if err != nil {
		handleInternalServerError(w, endpoint, method, queries, err)

		return
	}

	_, err = w.Write(encoded)
	if err != nil {
		handleInternalServerError(w, endpoint, method, queries, err)

		return
	}

	loggingSuccess(method, endpoint, http.StatusOK, queries)
}
