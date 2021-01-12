package errors

import "fmt"


type InvalidFileName struct {
	FileName string
}

func (err InvalidFileName) Error() string {
	return fmt.Sprintf("Invalid file name: %s", err.FileName)
}

type EmptyData struct {
}

func (err EmptyData) Error() string {
	return "Cannot save empty data"
}

type InvalidFilePath struct {
	FilePath string
}

func (err InvalidFilePath) Error() string {
	return fmt.Sprintf("Invalid path: %s", err.FilePath)
}
