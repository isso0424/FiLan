package server

import (
	"fmt"
	"net/http"
)

func queryNotEnoughError(
	w http.ResponseWriter,
	endpoint string,
	method string,
	query string,
) {
		errorMessage := fmt.Sprintf("Need query parameter: %s", query)
		handlerRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)

		return
}
