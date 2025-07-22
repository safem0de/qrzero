package entity

type GenerateStringRequest struct {
    Biller_id       string `json:"biller_id" binding:"required"`
    Account_num     string `json:"account_num" binding:"required"`
    Company_bank    string `json:"company_bank" binding:"required"`
    Amount          string `json:"amount" binding:"required"`
}