// qrzero\internal\03_infrastructure\custable.go

package infrastructure

import (
	"context"
	"database/sql"
	"qrzero/internal/01_entity"
	"qrzero/internal/02_application"
)

type custableRepository struct {
	db *sql.DB
	queries *Queries
}

func NewCustableRepository(db *sql.DB, queries *Queries) application.CustableService {
	return &custableRepository{db: db, queries: queries}
}

func (r *custableRepository) GetRecentActiveCustomers(ctx context.Context) ([]entity.Custable, error) {
	query := r.queries.GetRecentActiveCustomers

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entity.Custable
	for rows.Next() {
		var c entity.Custable
		err := rows.Scan(
			&c.BillerID,
			&c.AccountNum,
			&c.CompanyBank,
			&c.Name,
			&c.CustomerStatus,
			&c.CreatedDateTime,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, rows.Err()
}
