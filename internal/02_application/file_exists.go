package application

import (
    "qrzero/internal/01_entity"
)

type FileExistService interface {
    CheckFileExist(req entity.FileExistRequest) (entity.FileExistResponse, error)
}