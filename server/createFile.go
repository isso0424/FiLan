package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func createFileHandler(w http.ResponseWriter, r *http.Request) {
	type Query struct {
		Path string `schema:"path,required"`
		Name string `schema:"name,required"`
	}

	const endpoint = "/file"
	const method = "POST"
	query := Query{}

	err := decoder.Decode(&query, r.URL.Query())
	queries := []loggingQuery{{key: "path", value: query.Path}, {key: "name", value: query.Name}}
	if err != nil {
		handleRequestError(w, endpoint, method, http.StatusBadRequest, queries, fmt.Sprintf(notEnoughQuery, "name and path"))

		return
	}

	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleInternalServerError(w, endpoint, method, queries, err)

		return
	}

	file, err := controller.SaveFile(buffer, query.Name, query.Path)
	if err != nil {
		handleInternalServerError(w, endpoint, method, queries, err)

		return
	}

	err = domainWritebackToClient(file, w)
	if err != nil {
		handleInternalServerError(w, endpoint, method, queries, err)

		return
	}

	loggingSuccess(method, endpoint, http.StatusCreated, queries)
}
