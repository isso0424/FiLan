package filerepository

import (
	"FiLan/repository/filesystem"
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
	FileSystemRepository filesystem.Repository
}

// New is FileRepository constructor
func New(fileName string, storageDir string, config *gorm.Config) (FileRepository, error) {
	db, err := gorm.Open(sqlite.Open(fileName), config)
	if err != nil {
		return FileRepository{}, nil
	}
	db.AutoMigrate(&fileModel{})

	fsRepo := filesystem.Repository{ StorageDir: storageDir }

	repo := FileRepository{ DB: db, FileSystemRepository: fsRepo }

	return repo, nil
}
