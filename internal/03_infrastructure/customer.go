// internal\03_infrastructure\customer.go

package infrastructure

import (
	"encoding/json"
	"os"
	"qrzero/internal/01_entity"
	"qrzero/internal/02_application"
)

type customerServiceImpl struct {
	rabbit RabbitMQClient
}

func (s *customerServiceImpl) CheckFileExist(filename string) bool {
    _, err := os.Stat("./qrcode/" + filename)
    return !os.IsNotExist(err)
}

func (s *customerServiceImpl) PublishQRJob(req entity.QRJobRequest) error {
    b, _ := json.Marshal(req)
    return s.rabbit.Publish("qr_job", b)
}

func NewCustomerService(rabbit RabbitMQClient) application.CustomerService {
    return &customerServiceImpl{rabbit: rabbit}
}
