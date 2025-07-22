package entity

type GenerateQRRequest struct {
    QRString string `json:"qr" binding:"required"`
    Path     string `json:"path" binding:"required"`
}