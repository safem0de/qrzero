package application

import (
    "qrzero/internal/01_entity"
)

type FileService interface {
    ListFiles(path string) ([]entity.FileInfo, error)
}

