package server

import "net/http"

func deleteFileHandler(w http.ResponseWriter, r *http.Request) {
	type Query struct {
		Name string `schema:"name,required"`
		Path string `schema:"path,required"`
	}

	const endpoint = "file"
	const method = "DELETE"
	query := Query{}

	err := decoder.Decode(&query, r.URL.Query())
	if err != nil {
		queryNotEnoughError(w, endpoint, method, "name or path")

		return
	}

	file, err := controller.DeleteFile(query.Name, query.Path)
	if err != nil {
		errorMessage := err.Error()
		handlerRequestError(w, endpoint, method, http.StatusBadRequest, errorMessage)

		return
	}

	domainWritebackToClient(file, w, endpoint, method)
}
