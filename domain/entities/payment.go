package entities

import (
	"cesjb/types_"
)

type Payment struct {
	ID          int             `json:"id"`
	AssociateID int             `json:"associateID"`
	Competence  types_.DateOnly `json:"competence"`  // Mês/ano da mensalidade
	PaymentDate types_.DateOnly `json:"paymentDate"` // Dia que pagou
	Value       float64         `json:"value"`       // valor em centavos
	Status      bool            `json:"status"`
}
