// qrzero\internal\01_entity\custable.go
package entity

import "time"

type Custable struct {
	BillerID        string    `json:"biller_id"`
	AccountNum      string    `json:"account_num"`
	CompanyBank     string    `json:"company_bank"`
	Name            string    `json:"name"`
	CustomerStatus  int       `json:"customer_status"`
	CreatedDateTime time.Time `json:"created_datetime"`
}
