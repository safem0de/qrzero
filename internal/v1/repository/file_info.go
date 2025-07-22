// internal\v1\repository\file_info.go

package repository

import "time"

type FileInfo struct {
    Name    string    `json:"name"`
    ModTime time.Time `json:"mod_time"`
}
