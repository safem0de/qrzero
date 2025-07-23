package infrastructure

import (
	"os"
    "sort"
	"qrzero/internal/01_entity"
	"qrzero/internal/02_application"
)

type fileCheckingRepository struct{}

func NewFileCheckingRepository() application.FileCheckingService {
    return &fileCheckingRepository{}
}

func (s *fileCheckingRepository) ListFiles(path string) ([]entity.FileInfo, error) {
    entries, err := os.ReadDir(path)
    if err != nil {
        return nil, err
    }
    var files []entity.FileInfo
    for _, entry := range entries {
        if entry.Type().IsRegular() {
            info, err := entry.Info()
            if err != nil {
                continue
            }
            files = append(files, entity.FileInfo{
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
