package main

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"qrzero/internal/03_infrastructure"
	handler_v1 "qrzero/internal/04_api/v1/handler"
	handler_v2 "qrzero/internal/04_api/v2/handler"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	queries, err := infrastructure.LoadQueriesFromFile("configs/query.json")
	if err != nil {
		log.Fatal("cannot load queries: ", err)
	}

	rabbitClient, err := infrastructure.NewRabbitMQClient(os.Getenv("AMQP_URL"))
    if err != nil {
        log.Fatal("RabbitMQ error:", err)
    }

    customerService := infrastructure.NewCustomerService(rabbitClient)
	customerHandler := handler_v1.NewCustomerHandler(customerService)

	custableRepo := infrastructure.NewCustableRepository(db, queries)
	custableHandler := handler_v1.NewCustableHandler(custableRepo)

	fileCheckingRepo := infrastructure.NewFileCheckingRepository()
	fileCheckingHandler := handler_v1.NewFileHandler(fileCheckingRepo)

	fileExistRepo := infrastructure.NewFileExistRepository()
	fileExistHandler := handler_v1.NewFileExistHandler(fileExistRepo)

	genStringRepo := infrastructure.NewGenerateStringRepository()
	genStringHandler := handler_v1.NewGenerateStringHandler(genStringRepo)

	genQRRepo := infrastructure.NewQRRepository()
	genQRHandler := handler_v1.NewQRHandler(genQRRepo)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", handler_v1.HelloHandler)
		v1.GET("/custable", custableHandler.GetRecentActiveCustomers)
		v1.GET("/files", fileCheckingHandler.ListFiles)
		v1.GET("/file-exist", fileExistHandler.CheckFileExist)
		v1.POST("/generate", genStringHandler.GenerateString)
		v1.POST("/qr", genQRHandler.GenerateQR)
		v1.POST("/generate-qr-job", customerHandler.GenerateQRJob)
	}

	v2 := r.Group("/api/v2")
	{
		v2.GET("/hello", handler_v2.HelloHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3333")
}
