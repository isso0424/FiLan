package mockfile

import "FiLan/domain"

// FileRepository is mock struct for FileRepository
type FileRepository struct {
	Files []domain.File
}

// New is constructor for Mock FileRepository
func New() FileRepository {
	return FileRepository{ Files: []domain.File{} }
}

// Save is file save function
func (repo FileRepository) Save(file domain.File) {
	repo.Files = append(repo.Files, file)
}
