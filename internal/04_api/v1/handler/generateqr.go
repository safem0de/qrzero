package handler

import (
    "net/http"
    "qrzero/internal/01_entity"
    "qrzero/internal/02_application"
    "github.com/gin-gonic/gin"
)

type QRHandler struct {
    svc application.QRService
}

func NewQRHandler(svc application.QRService) *QRHandler {
    return &QRHandler{svc: svc}
}

// @Summary      Generate QR Code
// @Description  สร้างไฟล์ QR Code PNG จาก text
// @Tags         v1-POST
// @Accept       json
// @Produce      json
// @Param        body  body  entity.GenerateQRRequest  true  "ข้อมูล QR"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /api/v1/qr [post]
func (h *QRHandler) GenerateQR(c *gin.Context) {
    var req entity.GenerateQRRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.svc.GenerateQR(req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "QR code generated",
        "file":    req.Path,
    })
}
