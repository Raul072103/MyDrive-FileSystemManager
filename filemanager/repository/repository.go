package repository

import "os"

type FileManagerRepo interface {
	CreateFile(path string) error

	DeleteFile(path string) error

	UpdateFile(path string, file os.File) error

	ReadFile(path string) ([]byte, error)
}
