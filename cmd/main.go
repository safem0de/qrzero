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

	customerRepo := infrastructure.NewCustomerRepository(db, queries)
	customerHandler := handler_v1.NewCustomerHandler(customerRepo)

	fileRepo := infrastructure.NewFileRepository()
	fileHandler := handler_v1.NewFileHandler(fileRepo)

	genStringRepo := infrastructure.NewGenerateStringRepository()
	genStringHandler := handler_v1.NewGenerateStringHandler(genStringRepo)

	genQRRepo := infrastructure.NewQRRepository()
	genQRHandler := handler_v1.NewQRHandler(genQRRepo)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", handler_v1.HelloHandler)
		v1.GET("/customers", customerHandler.GetRecentActiveCustomers)
		v1.GET("/files", fileHandler.ListFiles)
		v1.POST("/generate", genStringHandler.GenerateString)
		v1.POST("/qr", genQRHandler.GenerateQR)
	}

	v2 := r.Group("/api/v2")
	{
		v2.GET("/hello", handler_v2.HelloHandler)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":3333")
}
