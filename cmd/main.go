// cmd\main.go

package main

import (
    "github.com/gin-gonic/gin"
    handler_v1 "qrzero/internal/v1/handler"
    handler_v2 "qrzero/internal/v2/handler"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    _ "qrzero/docs"
)

func main() {
    r := gin.Default()

    v1 := r.Group("/api/v1")
    {
        v1.GET("/hello", handler_v1.HelloHandler)
    }

    v2 := r.Group("/api/v2")
    {
        v2.GET("/hello", handler_v2.HelloHandler)
    }

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":3333")
}
