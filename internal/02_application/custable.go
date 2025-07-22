// qrzero\internal\02_application\custable.go

package application

import (
	"context"
	"qrzero/internal/01_entity"
)

type CustomerService interface {
	GetRecentActiveCustomers(ctx context.Context) ([]entity.Customer, error)
}