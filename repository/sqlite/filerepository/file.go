package filerepository

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type fileModel struct {
	gorm.Model
	Name string
	Path string
	Size int

	AddedAt time.Time
	RefreshedAt time.Time
}

// FileRepository is repository of file with gorm and sqlite
type FileRepository struct {
	DB *gorm.DB
	StorageDir string
}

// New is FileRepository constructor
func New(fileName string, storageDir string, config *gorm.Config) (FileRepository, error) {
	db, err := gorm.Open(sqlite.Open(fileName), config)
	if err != nil {
		return FileRepository{}, nil
	}
	db.AutoMigrate(&fileModel{})

	repo := FileRepository{ DB: db, StorageDir: storageDir }

	return repo, nil
}
