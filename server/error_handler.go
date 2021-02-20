package server

import (
	"errors"
	"log"
	"net/http"
)

const (
	notEnoughQuery = "Required query parameters: %s"
	failedParsingJSON = "Error occur in json parsing"
)

func handleRequestError(w http.ResponseWriter, endpoint string, method string, statusCode int, params []loggingQuery, errorMessage string) {
	loggingFailed(method, endpoint, http.StatusInternalServerError, params, errors.New(errorMessage))

	writeError(w, statusCode, errorMessage)
}

func handleInternalServerError(w http.ResponseWriter, endpoint string, method string, params []loggingQuery, occurredErr error) {
	log.Printf(logFormat, method, endpoint, http.StatusInternalServerError, occurredErr)
	loggingFailed(method, endpoint, http.StatusInternalServerError, params, occurredErr)

	writeError(w, http.StatusInternalServerError, "internal server error")
}
