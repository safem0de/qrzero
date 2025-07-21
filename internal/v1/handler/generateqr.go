package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "qrzero/internal/v1/service"
)

type QRRequest struct {
    QRString string `json:"qr" binding:"required"`
    Path     string `json:"path" binding:"required"`
}

type QRHandler struct {
    svc service.QRService
}

func NewQRHandler(svc service.QRService) *QRHandler {
    return &QRHandler{svc: svc}
}

// @Summary      Generate QR Code
// @Description  สร้างไฟล์ QR Code PNG จาก text
// @Tags         v1
// @Accept       json
// @Produce      json
// @Param        body  body  QRRequest  true  "ข้อมูล QR"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Router       /api/v1/qr [post]
func (h *QRHandler) GenerateQR(c *gin.Context) {
    var req QRRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := h.svc.GenerateQR(req.QRString, req.Path); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "message": "QR code generated",
        "file":    req.Path,
    })
}
