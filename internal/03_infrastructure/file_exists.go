package infrastructure

import (
    "os"
    "qrzero/internal/01_entity"
    "qrzero/internal/02_application"
)

type fileExistRepository struct{}

func NewFileExistRepository() application.FileExistService {
    return &fileExistRepository{}
}

func (s *fileExistRepository) CheckFileExist(req entity.FileExistRequest) (entity.FileExistResponse, error) {
    _, err := os.Stat(req.Path)
    if err == nil {
        return entity.FileExistResponse{Exists: true}, nil
    }
    if os.IsNotExist(err) {
        return entity.FileExistResponse{Exists: false}, nil
    }
    return entity.FileExistResponse{}, err // other error
}
