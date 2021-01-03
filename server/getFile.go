package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func getFileHandler(w http.ResponseWriter, r *http.Request) {
	const endpoint = "/file"
	var name string
	var path string
	err := decoder.Decode(name, r.URL.Query())
	if err != nil {
		errorMessage := "Need query parameter: name"
		log.Printf(logFormat, endpoint, http.StatusBadRequest, errorMessage)

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(errorMessage))
		if err != nil {
			log.Println(err)
		}

		return
	}
	err = decoder.Decode(path, r.URL.Query())
	if err != nil {
		errorMessage := "Need query parameter: path"
		log.Printf(logFormat, endpoint, http.StatusBadRequest, errorMessage)

		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(errorMessage))
		if err != nil {
			log.Println(err)
		}

		return
	}

	file, err := controller.GetFile(name, path)
	if err != nil {
		errorMessage := "Not found"
		log.Printf(logFormat, endpoint, http.StatusNotFound, errorMessage)
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write([]byte(errorMessage))
		if err != nil {
			log.Println(err)
		}

		return
	}

	model := convertDomainToModel(file)
	encoded, err := json.Marshal(model)
	if err != nil {
		errorMessage := "Not found"
		log.Printf(logFormat, endpoint, http.StatusInternalServerError, errorMessage)
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Error occur pasing json"))
		if err != nil {
			log.Println(err)
		}
	}

	_, err = w.Write(encoded)
	if err != nil {
		log.Println(err)
	}
}
