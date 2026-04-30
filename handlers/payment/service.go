package payment

import (
	"cesjb/domain/entities"
	"cesjb/types_"
)

type Service interface {
	CreatePayment(associateID int, competence, paymentDate types_.DateOnly, value float64,
		status bool) (entities.Payment, error)
	GetMonthlyTotal(competence types_.DateOnly) (float64, error)
	GetPaymentsByMonth(competence types_.DateOnly) ([]entities.PaymentWithAssociate, error)
	GetDefaultersByMonth(competence types_.DateOnly) ([]entities.Associate, error)
	GetPaymentsByAssociate(associateID int) ([]entities.Payment, error)
	UpdatePayment(id int, competence, paymentDate types_.DateOnly, value float64, status bool) (entities.Payment, error)
}
