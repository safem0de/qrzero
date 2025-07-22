// cmd\main.go

package main

import (
    "database/sql"
	"log"
	"os"
    _ "github.com/denisenkom/go-mssqldb"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    handler_v1 "qrzero/internal/v1/handler"
    service_v1 "qrzero/internal/v1/service"
    repository_v1 "qrzero/internal/v1/repository"
    handler_v2 "qrzero/internal/v2/handler"

    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
    _ "qrzero/docs"
)

func main() {
    // STEP 1: Load env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using OS environment only")
    }

    // STEP 2: Connect MSSQL
    db, err := sql.Open("sqlserver", os.Getenv("MSSQL_CONN"))
	if err != nil {
		log.Fatal("cannot connect to mssql:", err)
	}
	defer db.Close()

    r := gin.Default()

    genService := service_v1.NewGenerateService()
    genHandler := handler_v1.NewGenerateHandler(genService)

    qrService := service_v1.NewQRService()
    qrHandler := handler_v1.NewQRHandler(qrService)

    fileSvc := service_v1.NewFileService()
    fileHandler := handler_v1.NewFileHandler(fileSvc)

    // ===== Customer Repository/Service/Handler DI =====
    customerRepo := repository_v1.NewCustomerRepository(db)
    customerService := service_v1.NewCustomerService(customerRepo)
    customerHandler := handler_v1.NewCustomerHandler(customerService)

    v1 := r.Group("/api/v1")
    {
        v1.GET("/hello", handler_v1.HelloHandler)
        v1.GET("/files", fileHandler.ListFiles)
        v1.GET("/customers", customerHandler.GetRecentActiveCustomers)
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
