package server

import (
	"net/http"
	"strconv"
)

func createFileHandler(w http.ResponseWriter, r *http.Request) {
	type Query struct {
		Path string `schema:"path,required"`
		Size string `schema:"size,required"`
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

	bufferSize, err := strconv.Atoi(query.Size)
	if err != nil {
		errorMessage := "Query parameter size must be integer"
		handlerRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)

		return
	}

	buffer := make([]byte, bufferSize)
	_, err = r.Body.Read(buffer)
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
