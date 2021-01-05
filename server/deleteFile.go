package server

import "net/http"

func deleteFileHandler(w http.ResponseWriter, r *http.Request) {
	const endpoint = "file"
	const method = "DELETE"
	var name string
	var path string
	err := decoder.Decode(name, r.URL.Query())
	if err != nil {
		queryNotEnoughError(w, endpoint, method, "name")

		return
	}

	err = decoder.Decode(path, r.URL.Query())
	if err != nil {
		queryNotEnoughError(w, endpoint, method, "path")

		return
	}

	file, err := controller.DeleteFile(name, path)
	if err != nil {
		errorMessage := err.Error()
		handlerRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)

		return
	}

	domainWritebackToClient(file, w, endpoint, method)
}
