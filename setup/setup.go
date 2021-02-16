// Package setup provides setupper before serve
package setup

import (
	"FiLan/repository"
	"FiLan/repository/filesystem"
	"FiLan/repository/mock"
	"FiLan/repository/mock/accessrepository"
	"FiLan/repository/sqlite/filerepository"

	"gorm.io/gorm"
)

type setuppedObj struct {
	DbRepository repository.FileRepository
	FsRepository repository.FileAccessRepository
	Err error
}

func Setup(mode string, storageDir string, dbFile string) setuppedObj {
	if mode == "test" {
		mockDB := mock.New()
		mockFS := accessrepository.New()
		return setuppedObj{
			DbRepository: mockDB,
			FsRepository: &mockFS,
			Err: nil,
		}
	}

	err := fsSetup(storageDir)
	if err != nil {
		return setuppedObj{
			Err: err,
		}
	}

	fsRepository := filesystem.New(storageDir)
	dbRepository, err := filerepository.New(dbFile, &gorm.Config{})
	if err != nil {
		return setuppedObj{
			Err: err,
		}
	}

	return setuppedObj{
		DbRepository: dbRepository,
		FsRepository: fsRepository,
		Err: nil,
	}
}
