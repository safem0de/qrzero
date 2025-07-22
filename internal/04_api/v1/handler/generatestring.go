package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
    "qrzero/internal/01_entity"
	"qrzero/internal/02_application"
)

type GenerateStringHandler struct {
    svc application.GenerateStringService
}

func NewGenerateStringHandler(svc application.GenerateStringService) *GenerateStringHandler {
    return &GenerateStringHandler{svc: svc}
}

// @Summary      Generate String
// @Description  Generate string from input params
// @Tags         v1-POST
// @Accept       json
// @Produce      plain
// @Param        body  body  entity.GenerateStringRequest  true  "Request body"
// @Success      200   {string}  string
// @Router       /api/v1/generate [post]
func (h *GenerateStringHandler) GenerateString(c *gin.Context) {
    var req entity.GenerateStringRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := h.svc.GenerateString(req)
    c.String(http.StatusOK, result)
}
