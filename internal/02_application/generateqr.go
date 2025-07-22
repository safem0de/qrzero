package application

import "qrzero/internal/01_entity"

type QRService interface {
    GenerateQR(req entity.GenerateQRRequest) error
}