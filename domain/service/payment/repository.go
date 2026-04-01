package payment

import (
	"cesjb/domain/entities"
)

type Repository interface {
	SavePayment(payment entities.Payment) (entities.Payment, error)
}
