package service

import "fmt"

type GenerateService interface {
    Generate(accountID, groupID, customerID string) string
}

type generateService struct{}

func NewGenerateService() GenerateService {
    return &generateService{}
}

func (s *generateService) Generate(accountID, groupID, customerID string) string {
    // Generate output as requested (each line with \r\n, last line is 0)
    return fmt.Sprintf("|%s\r\n%s%s\r\n\r\n0", accountID, groupID, customerID)
}
