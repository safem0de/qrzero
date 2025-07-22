package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "qrzero/internal/v1/service"
)

type GenerateRequest struct {
    AccountID  string `json:"account_id" binding:"required"`
    GroupID    string `json:"group_id" binding:"required"`
    CustomerID string `json:"customer_id"`
}

type GenerateHandler struct {
    svc service.GenerateService
}

func NewGenerateHandler(svc service.GenerateService) *GenerateHandler {
    return &GenerateHandler{svc: svc}
}

// @Summary      Generate String
// @Description  Generate string from input params
// @Tags         v1-POST
// @Accept       json
// @Produce      plain
// @Param        body  body  GenerateRequest  true  "Request body"
// @Success      200   {string}  string
// @Router       /api/v1/generate [post]
func (h *GenerateHandler) Generate(c *gin.Context) {
    var req GenerateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result := h.svc.Generate(req.AccountID, req.GroupID, req.CustomerID)
    c.String(http.StatusOK, result)
}
