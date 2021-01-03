// Package filesystem provides control to filesystem
package filesystem

// Repository is control filesystem repository
type Repository struct {
	StorageDir string
}

// New is constructor for repository filesystem
func New(storageDir string) Repository {
	return Repository{ StorageDir: storageDir }
}
