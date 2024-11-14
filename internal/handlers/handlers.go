package handlers

import (
	"MyDrive-FileSystemManager/internal/repository"
)

// Repo the repository used by the handlers
var Repo *repository.Repository

// NewHandlers sets the repository for the handlers
func NewHandlers(r *repository.Repository) {
	Repo = r
}
