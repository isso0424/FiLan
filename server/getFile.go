package server

import (
	"fmt"
	"net/http"
)

func getFileHandler(w http.ResponseWriter, r *http.Request) {
	type Query struct {
		Name string `schema:"name,required"`
		Path string `schema:"path,required"`
	}

	const endpoint = "/file"
	const method = "GET"
	query := Query{}

	err := decoder.Decode(&query, r.URL.Query())
	queries := []loggingQuery{{key: "path", value: query.Path}, {key: "name", value: query.Name}}
	if err != nil {
		handleRequestError(w, endpoint, method, http.StatusBadRequest, queries, fmt.Sprintf(notEnoughQuery, "name and path"))

		return
	}

	file, err := controller.GetFile(query.Name, query.Path)
	if err != nil {
		errorMessage := "Not found"
		handleRequestError(w, endpoint, method, http.StatusNotFound, queries, errorMessage)

		return
	}

	err = fileWritebackToClient(file.Data, w)
	if err != nil {
		handleInternalServerError(w, endpoint, method, queries, err)

		return
	}

	loggingSuccess(method, endpoint, http.StatusOK, queries)
}
