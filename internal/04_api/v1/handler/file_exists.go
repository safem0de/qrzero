package handler

import (
    "net/http"
    "qrzero/internal/01_entity"
    "qrzero/internal/02_application"
    "github.com/gin-gonic/gin"
)

type FileExistHandler struct {
    svc application.FileExistService
}

func NewFileExistHandler(svc application.FileExistService) *FileExistHandler {
    return &FileExistHandler{svc: svc}
}

// @Summary      Check file exist
// @Description  ตรวจสอบว่าไฟล์ path นี้มีอยู่จริงไหม
// @Tags         v1-GET
// @Accept       json
// @Produce      json
// @Param        path query string true "File path"
// @Success      200  {object}  entity.FileExistResponse
// @Failure      400  {object}  map[string]string
// @Router       /api/v1/file-exist [get]
func (h *FileExistHandler) CheckFileExist(c *gin.Context) {
    path := c.Query("path")
    if path == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "path is required"})
        return
    }
    req := entity.FileExistRequest{Path: path}
    resp, err := h.svc.CheckFileExist(req)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, resp)
}
