// internal\04_api\v1\handler\hello.go

package handler

import (
	"github.com/gin-gonic/gin"
)

// HelloHandler godoc
// @Summary      Say Hello (V1)
// @Description  Hello from v1
// @Tags         v1-GET
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /api/v1/hello [get]
func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello from v1"})
}
