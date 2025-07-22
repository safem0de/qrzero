// qrzero\internal\01_entity\custable.go
package entity

import "time"

type Customer struct {
	BillerID        string    `json:"biller_id"`
	AccountNum      string    `json:"account_num"`
	CompanyBank     string    `json:"company_bank"`
	Name            string    `json:"name"`
	CustomerStatus  int       `json:"customer_status"`
	ModifiedDateTime time.Time `json:"modified_datetime"`
}
