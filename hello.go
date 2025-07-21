package main

import (
    "github.com/gin-gonic/gin"
)

// HelloHandler godoc
// @Summary      Say Hello
// @Description  Respond with hello world
// @Tags         example
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /hello [get]

func HelloHandler(c *gin.Context) {
    c.JSON(200, gin.H{"message": "hello world"})
}
