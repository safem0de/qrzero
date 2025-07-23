// internal\04_api\v1\handler\customer.go

package handler

import (
    "qrzero/internal/01_entity"
    "qrzero/internal/02_application"
    "github.com/gin-gonic/gin"
)

type CustomerHandler struct {
    svc application.CustomerService
}

// *** Constructor ***
func NewCustomerHandler(svc application.CustomerService) *CustomerHandler {
    return &CustomerHandler{svc: svc}
}

// GenerateQRJob godoc
// @Summary      Queue QR Job Generation
// @Description  ส่งข้อมูลไป Queue เพื่อ generate QR (RabbitMQ/Async)
// @Tags         v1-POST
// @Accept       json
// @Produce      json
// @Param        body body entity.QRJobRequest true "QR job request data"
// @Success      202 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /api/v1/generate-qr-job [post]
func (h *CustomerHandler) GenerateQRJob(c *gin.Context) {
    var req entity.QRJobRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    if h.svc.CheckFileExist(req.AccountNum + ".png") {
        c.JSON(200, gin.H{"status":"exists", "message":"File already exists"})
        return
    }

    if err := h.svc.PublishQRJob(req); err != nil {
        c.JSON(500, gin.H{"status":"error", "message": err.Error()})
        return
    }

    c.JSON(202, gin.H{"status": "queued", "message": "QR Generation job queued"})
}
