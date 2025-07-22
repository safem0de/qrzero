package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "qrzero/internal/v1/service"
)

type FileHandler struct {
    svc service.FileService
}

func NewFileHandler(svc service.FileService) *FileHandler {
    return &FileHandler{svc: svc}
}

// @Summary      List files in directory
// @Description  Show filename and last modified time in a directory
// @Tags         v1-GET
// @Accept       json
// @Produce      json
// @Param        path query string true "Directory path"
// @Success      200  {array}  repository.FileInfo
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/files [get]
func (h *FileHandler) ListFiles(c *gin.Context) {
    path := c.Query("path")
    if path == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "path is required"})
        return
    }
    files, err := h.svc.ListFiles(path)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, files)
}
