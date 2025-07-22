// cmd\main.go

package main

import (
    "github.com/gin-gonic/gin"
    handler_v1 "qrzero/internal/v1/handler"
    service_v1 "qrzero/internal/v1/service"
    handler_v2 "qrzero/internal/v2/handler"
    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    _ "qrzero/docs"
)

func main() {
    r := gin.Default()

    genService := service_v1.NewGenerateService()
    genHandler := handler_v1.NewGenerateHandler(genService)

    qrService := service_v1.NewQRService()
    qrHandler := handler_v1.NewQRHandler(qrService)

    fileSvc := service_v1.NewFileService()
    fileHandler := handler_v1.NewFileHandler(fileSvc)

    v1 := r.Group("/api/v1")
    {
        v1.GET("/hello", handler_v1.HelloHandler)
        v1.GET("/files", fileHandler.ListFiles)
        v1.POST("/generate", genHandler.Generate)
        v1.POST("/qr", qrHandler.GenerateQR)
    }

    v2 := r.Group("/api/v2")
    {
        v2.GET("/hello", handler_v2.HelloHandler)
    }

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":3333")
}
