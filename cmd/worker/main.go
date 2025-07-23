// cmd\worker\main.go

package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/streadway/amqp"
    "qrzero/internal/01_entity"
    "github.com/skip2/go-qrcode"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, using OS env only")
    }
	
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

    msgs, err := ch.Consume(
        "qr_job", // queue
        "",       // consumer
        true,     // auto-ack (หรือใช้ manual ถ้าต้องการ)
        false,    // exclusive
        false,    // no-local
        false,    // no-wait
        nil,      // args
    )
    if err != nil {
        log.Fatalf("Queue consume error: %v", err)
    }

    log.Println(" [*] Waiting for QR jobs. To exit press CTRL+C")

    for msg := range msgs {
        var job entity.QRJobRequest
        if err := json.Unmarshal(msg.Body, &job); err != nil {
            log.Println("Bad message:", err)
            continue
        }
        
        fmt.Printf("Received job: %+v\n", job)

        qrtext := fmt.Sprintf("|%s\n%s%s\n\n%s", job.BillerID, job.AccountNum, job.CompanyBank, job.Amount)

        savePath := "./qrcode/" + job.FilePath
        if err := qrcode.WriteFile(qrtext, qrcode.Medium, 256, savePath); err != nil {
            log.Printf("QR Generate failed: %v", err)
        } else {
            log.Printf("QR Generated: %s", savePath)
        }
    }
}
