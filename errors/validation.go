// Package errors provides some errors
package errors

import "fmt"

// InvalidFileName is error type for invalid file name
type InvalidFileName struct {
	FileName string
}

func (err InvalidFileName) Error() string {
	return fmt.Sprintf("Invalid file name: %s", err.FileName)
}

// EmptyData is error type for saving empty data
type EmptyData struct {
}

func (err EmptyData) Error() string {
	return "Cannot save empty data"
}

// InvalidFilePath is error type for invalid file path
type InvalidFilePath struct {
	FilePath string
}

func (err InvalidFilePath) Error() string {
	return fmt.Sprintf("Invalid path: %s", err.FilePath)
}
