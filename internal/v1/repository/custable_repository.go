package repository

import (
	"context"
	"database/sql"
)

type CustomerRepository interface {
	GetRecentActiveCustomers(ctx context.Context) ([]Customer, error)
}

type customerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) GetRecentActiveCustomers(ctx context.Context) ([]Customer, error) {
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

	var customers []Customer
	for rows.Next() {
		var c Customer
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
