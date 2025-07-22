// internal\v1\service\custable_service.go

package service

import (
	"context"
	"qrzero/internal/v1/repository"
)

type CustomerService interface {
	GetRecentActiveCustomers(ctx context.Context) ([]repository.Customer, error)
}

type customerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerService{repo: repo}
}

func (s *customerService) GetRecentActiveCustomers(ctx context.Context) ([]repository.Customer, error) {
	return s.repo.GetRecentActiveCustomers(ctx)
}