// internal\01_entity\customer.go

package entity

type QRJobRequest struct {
    BillerID     string `json:"biller_id"`
    AccountNum   string `json:"account_num"`
    CompanyBank  string `json:"company_bank"`
    Amount       string `json:"amount"`
    FilePath     string `json:"file_path"`
    QRString     string `json:"qr_string"`
}
