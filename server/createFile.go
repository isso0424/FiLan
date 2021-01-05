package server

import (
	"net/http"
	"strconv"
)

func createFileHandler(w http.ResponseWriter, r *http.Request) {
	const endpoint = "/file"
	const method = "POST"
	var name string
	var path string
	var size string
	err := decoder.Decode(name, r.URL.Query())
	if err != nil {
		errorMessage := "Need query parameter: name"
		handlerRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)

		return
	}
	err = decoder.Decode(path, r.URL.Query())
	if err != nil {
		errorMessage := "Need query parameter: path"
		handlerRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)

		return
	}

	err = decoder.Decode(size, r.URL.Query())
	if err != nil {
		errorMessage := "Need query parameter: size"
		handlerRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)

		return
	}

	bufferSize, err := strconv.Atoi(size)
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

	file, err := controller.SaveFile(buffer, name, path)
	if err != nil {
		errorMessage := err.Error()
		handlerRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)

		return
	}

	domainWritebackToClient(file, w, endpoint, method)
}
