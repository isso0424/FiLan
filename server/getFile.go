package server

import (
	"log"
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
	if err != nil {
		handleInvalidQuery(w, endpoint, method, "name or path")

		return
	}

	file, err := controller.GetFile(query.Name, query.Path)
	if err != nil {
		errorMessage := "Not found"
		handleRequestError(w, endpoint, method, http.StatusNotFound, errorMessage)

		return
	}

	log.Println(len(file.Data))

	err = fileWritebackToClient(file.Data, w)
	if err != nil {
		handleInternalServerError(w, endpoint, method, err)

		return
	}

	loggingSuccess(method, endpoint, http.StatusOK)
}
