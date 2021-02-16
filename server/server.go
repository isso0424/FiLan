// Package server is server functions with gorilla
package server

import (
	"FiLan/controller/filer"
	"FiLan/repository/filesystem"
	"FiLan/repository/sqlite/filerepository"
	"FiLan/usecase"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"gorm.io/gorm"

	"net/http"
	"time"
)

const (
	logFormat       = "Endpoint: %s Status: %d Description: %s\n"
	timeoutDuration = 15
)

var (
	controller usecase.Filer
	decoder    = schema.NewDecoder()
)

// Serve is function lanching server
func Serve(dbfile string, storageDir string) error {
	accessRepository := filesystem.New(storageDir)
	repo, err := filerepository.New(dbfile, &gorm.Config{})
	if err != nil {
		return err
	}
	controller = filer.New(repo, accessRepository)

	router := mux.NewRouter()

	router.HandleFunc("/file", getFileHandler).Queries("name", "{name}", "path", "{path}").Methods("GET")
	router.
		HandleFunc("/file", createFileHandler).
		Queries("name", "{name}", "path", "{path}", "size", "{size}").
		Methods("POST")
	router.HandleFunc("/file", deleteFileHandler).Queries("name", "{name}", "path", "{path}").Methods("DELETE")

	router.HandleFunc("/files", getFilesHandler).Queries("path", "{path}").Methods("GET")

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: timeoutDuration * time.Second,
		ReadTimeout:  timeoutDuration * time.Second,
	}

	return server.ListenAndServe()
}
