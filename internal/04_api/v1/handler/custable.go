// internal\04_api\v1\handler\custable.go

package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"qrzero/internal/02_application"
)

// ให้ struct CustomerHandler รับ CustomerRepository (หรือ Service) จาก application
type CustableHandler struct {
	service application.CustableService
}

func NewCustableHandler(service application.CustableService) *CustableHandler {
	return &CustableHandler{service: service}
}

// GetRecentActiveCustomers godoc
// @Summary      รายชื่อลูกค้าแอคทีฟในสัปดาห์นี้
// @Description  คืนค่าข้อมูลลูกค้าที่แอคทีฟหรือ Re-Active ในสัปดาห์ปัจจุบัน (ตาม MSSQL)
// @Tags         v1-GET
// @Accept       json
// @Produce      json
// @Success      200  {array}  entity.Custable
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/custable [get]
func (h *CustableHandler) GetRecentActiveCustomers(c *gin.Context) {
	customers, err := h.service.GetRecentActiveCustomers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}
