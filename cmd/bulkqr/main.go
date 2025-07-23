package main

import (
	"context"
    "database/sql"
    "encoding/json"
    "log"
    "os"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
    "github.com/streadway/amqp"
	"qrzero/internal/01_entity"
	"qrzero/internal/03_infrastructure"
)

func main() {
	// 1. Load env
    godotenv.Load()

	// 2. RabbitMq
    amqpURL := os.Getenv("AMQP_URL")
    conn, err := amqp.Dial(amqpURL)
    if err != nil {
        log.Fatalf("RabbitMQ connect error: %v", err)
    }
    defer conn.Close()

    ch, err := conn.Channel()
    if err != nil {
        log.Fatalf("Channel error: %v", err)
    }
    defer ch.Close()

	// 3. Connect DB
	db, err := sql.Open("sqlserver", os.Getenv("MSSQL_CONN"))
    if err != nil {
        log.Fatalf("cannot connect to mssql: %v", err)
    }
    defer db.Close()

	// 4. Load queries
    queries, err := infrastructure.LoadQueriesFromFile("configs/query.json")
    if err != nil {
        log.Fatal("cannot load queries: ", err)
    }

	// 5. Get all active customers
    custRepo := infrastructure.NewCustableRepository(db, queries)
    customers, err := custRepo.GetRecentActiveCustomers(context.Background())
    if err != nil {
        log.Fatal("get customers error:", err)
    }

    for _, c := range customers {
        req := entity.QRJobRequest{
            BillerID:    c.BillerID,
            AccountNum:  c.AccountNum,
            CompanyBank: c.CompanyBank,
            Amount:      "0",
            FilePath:    c.AccountNum + ".png",
        }
        body, _ := json.Marshal(req)
        err := ch.Publish(
            "",         // exchange
            "qr_job",   // queue name
            false,      // mandatory
            false,      // immediate
            amqp.Publishing{
                ContentType: "application/json",
                Body:        body,
            })
        if err != nil {
            log.Println("Send error:", err)
        } else {
            log.Printf("Queued: %s", c.AccountNum)
        }
    }
    log.Println("All jobs sent!")
}