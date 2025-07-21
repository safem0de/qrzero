package main

import (
    "github.com/gin-gonic/gin"
    _ "qrzero/docs"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

// @title           Example API
// @version         1.0
// @description     This is a sample server.
// @host            localhost:8080
// @BasePath        /

func main() {
    r := gin.Default()
    r.GET("/hello", HelloHandler)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":3333")
}
