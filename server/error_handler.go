package server

import (
	"fmt"
	"log"
	"net/http"
)

func handleInvalidQuery(
	w http.ResponseWriter,
	endpoint string,
	method string,
	query string,
) {
	errorMessage := fmt.Sprintf("Required query parameters: %s\n", query)
	handleRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)
}

func handleJSONParseError(
	w http.ResponseWriter,
	endpoint string,
	method string,
) {
	errorMessage := "error occur in json parsing"
	handleRequestError(w, endpoint, method, http.StatusInternalServerError, errorMessage)
}

func handleRequestError(w http.ResponseWriter, endpoint string, method string, statusCode int, errorMessage string) {
	log.Printf(logFormat, method, endpoint, statusCode, errorMessage)

	writeError(w, statusCode, errorMessage)
}

func handleInternalServerError(w http.ResponseWriter, endpoint string, method string, occurredErr error) {
	log.Printf(logFormat, method, endpoint, http.StatusInternalServerError, occurredErr)

	writeError(w, http.StatusInternalServerError, "internal server error")
}
