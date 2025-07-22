// qrzero\internal\03_infrastructure\custable.go

package infrastructure

import (
	"context"
	"database/sql"
	"qrzero/internal/01_entity"
	"qrzero/internal/02_application"
)

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) application.CustomerService {
	return &customerRepository{db: db}
}

func (r *customerRepository) GetRecentActiveCustomers(ctx context.Context) ([]entity.Customer, error) {
	query := `
	SELECT [BTL_BILLERID], [ACCOUNTNUM], [BPC_COMPANYBANK], [NAME], [BTL_CUSTOMERSTATUS], [MODIFIEDDATETIME]
	FROM [WEBORDER].[dbo].[CUSTTABLE]
	WHERE [MODIFIEDDATETIME] >= DATEADD(WEEK, DATEDIFF(WEEK, 0, GETDATE()), 0)
	  AND [MODIFIEDDATETIME] <  DATEADD(WEEK, DATEDIFF(WEEK, 0, GETDATE()) + 1, 0)
	  AND [BTL_CUSTOMERSTATUS] IN (0, 2)
	  AND [BTL_BILLERID] != ''
	ORDER BY [MODIFIEDDATETIME] DESC;`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entity.Customer
	for rows.Next() {
		var c entity.Customer
		err := rows.Scan(
			&c.BillerID,
			&c.AccountNum,
			&c.CompanyBank,
			&c.Name,
			&c.CustomerStatus,
			&c.ModifiedDateTime,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, rows.Err()
}
