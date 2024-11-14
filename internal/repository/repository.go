package repository

import (
	"MyDrive-FileSystemManager/internal/config"
)

// Repository is the repository type
type Repository struct {
	App      *config.AppConfig
	FileRepo FileManagerRepo
}

// New creates a new repository
func New(a *config.AppConfig, fileRepo FileManagerRepo) *Repository {
	return &Repository{
		a,
		fileRepo,
	}
}

type FileManagerRepo interface {
	CreateFile(path string) (err error)
	DeleteFile(path string) (err error)
	UpdateFile(path string, content []byte, updateAt int64) (err error)
	ReadFile(path string) ([]byte, error)
}
