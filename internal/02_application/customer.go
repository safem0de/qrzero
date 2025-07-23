package application

import (
	"qrzero/internal/01_entity"
)

type CustomerService interface {
    CheckFileExist(filename string) bool
    PublishQRJob(req entity.QRJobRequest) error
}
