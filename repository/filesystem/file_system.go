// Package filesystem provides control to filesystem
package filesystem

import (
	"os"
	"path"
)

// Repository is control filesystem repository
type Repository struct {
	StorageDir string
}

// New is constructor for repository filesystem
func New(storageDir string) Repository {
	cacheDir := os.Getenv("XDG_CACHE_HOME")
	if cacheDir == "" {
		homeDir := os.Getenv("HOME")
		cacheDir = path.Join(homeDir, ".cache")
	}

	storagePath := path.Join(cacheDir, storageDir)
	return Repository{StorageDir: storagePath}
}
