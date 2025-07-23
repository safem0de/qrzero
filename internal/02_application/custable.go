// qrzero\internal\02_application\custable.go

package application

import (
	"context"
	"qrzero/internal/01_entity"
)

type CustableService interface {
	GetRecentActiveCustomers(ctx context.Context) ([]entity.Custable, error)
}