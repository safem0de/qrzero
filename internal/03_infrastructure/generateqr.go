package infrastructure

import (
    "qrzero/internal/01_entity"
    "qrzero/internal/02_application"
    "github.com/skip2/go-qrcode"
)

type qrRepository struct{}

func NewQRRepository() application.QRService {
    return &qrRepository{}
}

func (s *qrRepository) GenerateQR(req entity.GenerateQRRequest) error {
    // แทนที่จะเขียน text ธรรมดา ให้ใช้ go-qrcode generate ภาพ
    return qrcode.WriteFile(req.QRString, qrcode.Medium, 256, req.Path)
}
