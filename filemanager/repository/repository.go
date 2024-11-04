package repository

import "os"

type FileManagerRepo interface {
	CreateFile(path string) error

	DeleteFile(path string) error

	UpdateFile(path string, file os.File) error

	ReadFile(path string) ([]byte, error)
}

//type DirectoryManagerRepo interface {
//	CreateDirectory(path string) error
//
//	DeleteDirectory(path string) error
//
//	UpdateDirectory(path string, file os.File) error
//
//	ReadDirectoryFiles(path string) ([]string, error)
//}
