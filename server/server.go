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

	router.HandleFunc("/file", getFileHandler).Methods("GET")
	router.HandleFunc("/file", createFileHandler).Methods("POST")
	router.HandleFunc("/file", deleteFileHandler).Methods("DELETE")

	router.HandleFunc("/files", getFilesHandler).Methods("GET")

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: timeoutDuration * time.Second,
		ReadTimeout:  timeoutDuration * time.Second,
	}

	return server.ListenAndServe()
}
