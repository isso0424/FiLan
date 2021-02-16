package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func getFilesHandler(w http.ResponseWriter, r *http.Request) {
	const endpoint = "/files"
	const method = "GET"
	var path string
	err := decoder.Decode(path, r.URL.Query())
	if err != nil {
		queryNotEnoughError(w, endpoint, method, "path")

		return
	}

	files, err := controller.GetFiles(path)
	if err != nil {
		errorMessage := "Internal server error"
		handlerRequestError(w, endpoint, method, http.StatusInternalServerError, errorMessage)

		return
	}

	filesModel := []fileModel{}

	for _, file := range files {
		fileModel := convertDomainToModel(file)
		filesModel = append(filesModel, fileModel)
	}

	encoded, err := json.Marshal(filesModel)
	if err != nil {
		errorMessage := "error occur in json pasing"
		handlerRequestError(w, endpoint, method, http.StatusInternalServerError, errorMessage)

		return
	}

	_, err = w.Write(encoded)
	if err != nil {
		log.Println(err)
	}
}
