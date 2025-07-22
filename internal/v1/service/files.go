package service

import (
    "os"
    "sort"
    "qrzero/internal/v1/repository"
)

type FileService interface {
    ListFiles(path string) ([]repository.FileInfo, error)
}

type fileService struct{}

func NewFileService() FileService {
    return &fileService{}
}

func (s *fileService) ListFiles(path string) ([]repository.FileInfo, error) {
    entries, err := os.ReadDir(path)
    if err != nil {
        return nil, err
    }
    var files []repository.FileInfo
    for _, entry := range entries {
        if entry.Type().IsRegular() {
            info, err := entry.Info()
            if err != nil {
                continue
            }
            files = append(files, repository.FileInfo{
                Name:    entry.Name(),
                ModTime: info.ModTime(),
            })
        }
    }
    // Sort by ModTime descending (ล่าสุดอยู่บนสุด)
    sort.Slice(files, func(i, j int) bool {
        return files[i].ModTime.After(files[j].ModTime)
    })
    // Limit เฉพาะ 1000 ไฟล์
    if len(files) > 1000 {
        files = files[:1000]
    }
    return files, nil
}
