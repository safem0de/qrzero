package application

import "qrzero/internal/01_entity"

type GenerateStringService interface {
    GenerateString(req entity.GenerateStringRequest) string
}
