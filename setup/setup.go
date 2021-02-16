// Package setup provides setupper before serve
package setup

import (
	"FiLan/repository"
	"FiLan/repository/filesystem"
	"FiLan/repository/mock/db"
	mockfs "FiLan/repository/mock/filesystem"
	"FiLan/repository/sqlite/filerepository"

	"gorm.io/gorm"
)

// SetuppedObj is object for setupped values
type SetuppedObj struct {
	DBRepository repository.FileRepository
	FsRepository repository.FileAccessRepository
	Err          error
}

// Setup is setup function for serve
func Setup(mode string, storageDir string, dbFile string) SetuppedObj {
	if mode == "test" {
		mockDB := db.New()
		mockFS := mockfs.New()

		return SetuppedObj{
			DBRepository: mockDB,
			FsRepository: &mockFS,
			Err:          nil,
		}
	}

	err := fsSetup(storageDir)
	if err != nil {
		return SetuppedObj{
			Err: err,
		}
	}

	fsRepository := filesystem.New(storageDir)
	dbRepository, err := filerepository.New(dbFile, &gorm.Config{})
	if err != nil {
		return SetuppedObj{
			Err: err,
		}
	}

	return SetuppedObj{
		DBRepository: dbRepository,
		FsRepository: fsRepository,
		Err:          nil,
	}
}
