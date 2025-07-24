// internal\04_api\v2\handler\hello.go

package handler

import (
    "github.com/gin-gonic/gin"
)

// HelloHandler godoc
// @Summary      Say Hello (V2)
// @Description  Hello from v2 (new logic)
// @Tags         v2
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /api/v2/hello [get]
func HelloHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "hello from v2 (new logic)"})
}
