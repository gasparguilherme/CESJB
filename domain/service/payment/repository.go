package payment

import (
	"cesjb/domain/entities"
	"cesjb/types_"
)

type Repository interface {
	SavePayment(payment entities.Payment) (entities.Payment, error)
	GetMonthlyTotal(competence types_.DateOnly) (float64, error)
}
