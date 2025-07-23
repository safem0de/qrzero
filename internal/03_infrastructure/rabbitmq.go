// internal/03_infrastructure/rabbitmq.go

package infrastructure

import (
    "github.com/streadway/amqp"
) 

// Interface
type RabbitMQClient interface {
    Publish(queue string, body []byte) error
}

// Implementation
type rabbitMQClientImpl struct {
    conn    *amqp.Connection
    channel *amqp.Channel
}

// NewRabbitMQClient สร้าง Client ใหม่
func NewRabbitMQClient(amqpURL string) (RabbitMQClient, error) {
    conn, err := amqp.Dial(amqpURL)
    if err != nil {
        return nil, err
    }
    ch, err := conn.Channel()
    if err != nil {
        return nil, err
    }
    return &rabbitMQClientImpl{conn: conn, channel: ch}, nil
}

// Publish ส่ง message เข้า queue
func (r *rabbitMQClientImpl) Publish(queue string, body []byte) error {
    _, err := r.channel.QueueDeclare(queue, true, false, false, false, nil)
    if err != nil {
        return err
    }
    return r.channel.Publish(
        "", queue, false, false,
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
}
