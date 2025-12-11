package associate

import (
	"cesjb/types_"
)

type UpdateAssociate struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	Email           string          `json:"email"`
	Tel             string          `json:"tel"`
	DateOfBirth     types_.DateOnly `json:"date_of_birth"`    // data de nascimento
	AssociationDate types_.DateOnly `json:"association_date"` // data de associacao
	Address         string          `json:"address"`
	DonationValue   float64         `json:"donation_value"` // valor da doacao
	PaymentDate     types_.DateOnly `json:"payment_date"`   // ultima data de pagamento
	Status          bool            `json:"status"`
}
