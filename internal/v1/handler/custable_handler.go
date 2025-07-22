// internal\v1\handler\custable_handler.go

package handler

import (
	"net/http"
	"qrzero/internal/v1/service"
	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service service.CustomerService
}

func NewCustomerHandler(service service.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

// GetRecentActiveCustomers godoc
// @Summary      รายชื่อลูกค้าแอคทีฟในสัปดาห์นี้
// @Description  คืนค่าข้อมูลลูกค้าที่แอคทีฟหรือ Re-Active ในสัปดาห์ปัจจุบัน (ตาม MSSQL)
// @Tags         v1-GET
// @Accept       json
// @Produce      json
// @Success      200  {array}  repository.Customer
// @Failure      500  {object}  map[string]string
// @Router       /api/v1/customers [get]
func (h *CustomerHandler) GetRecentActiveCustomers(c *gin.Context) {
	customers, err := h.service.GetRecentActiveCustomers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}
