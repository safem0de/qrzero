package application

import (
    "qrzero/internal/01_entity"
)

type FileCheckingService interface {
    ListFiles(path string) ([]entity.FileInfo, error)
}

