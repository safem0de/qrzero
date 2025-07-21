package service

import (
    "github.com/skip2/go-qrcode"
)

type QRService interface {
    GenerateQR(content, path string) error
}

type qrService struct{}

func NewQRService() QRService {
    return &qrService{}
}

func (s *qrService) GenerateQR(content, path string) error {
    // สร้าง QR PNG และเซฟไปที่ path
    return qrcode.WriteFile(content, qrcode.Medium, 256, path)
}
