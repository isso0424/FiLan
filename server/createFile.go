package server

import (
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
	if err != nil {
		queryNotEnoughError(w, endpoint, method, "name, size, or path")

		return
	}

	buffer, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errorMessage := "Received buffer larger than size"
		handlerRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)

		return
	}

	file, err := controller.SaveFile(buffer, query.Name, query.Path)
	if err != nil {
		errorMessage := err.Error()
		handlerRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)

		return
	}

	domainWritebackToClient(file, w, endpoint, method)
}
