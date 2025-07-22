package infrastructure

import (
	"fmt"
	"qrzero/internal/01_entity"
	"qrzero/internal/02_application"
)

type generateStringRepository struct{}

func NewGenerateStringRepository() application.GenerateStringService {
	return &generateStringRepository{}
}

func (s *generateStringRepository) GenerateString(req entity.GenerateStringRequest) string {
	// Generate output as requested (each line with \r\n, last line is 0)
	return fmt.Sprintf("|%s\\n%s%s\\n\\n%s", req.Biller_id, req.Account_num, req.Company_bank, req.Amount)
}
