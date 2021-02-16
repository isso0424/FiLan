package setup

import (
	"fmt"
	"os"
	"path"
)

type invalidPath struct {
	description string
}

func (err invalidPath) Error() string {
	return err.description
}

func fsSetup(storageDir string) error {
	cacheDir := os.Getenv("XDG_CACHE_HOME")
	if cacheDir == "" {
		homeDir := os.Getenv("HOME")
		cacheDir = path.Join(homeDir, ".cache")
	}

	storagePath := path.Join(cacheDir, storageDir)
	fmt.Println(storagePath)

	f, err := os.Stat(storagePath)
	if err == nil {
		if f.IsDir() {
			return nil
		}

		return invalidPath{ description: "Dir is already exist as file" }
	}

	return os.MkdirAll(storagePath, 0777)
}
