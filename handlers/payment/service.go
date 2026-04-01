package payment

import (
	"cesjb/domain/entities"
	"cesjb/types_"
)

type Service interface {
	CreatePayment(associateID int, competence, paymentDate types_.DateOnly, value float64,
		status bool) (entities.Payment, error)
}
