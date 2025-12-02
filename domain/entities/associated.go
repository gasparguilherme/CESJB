package entities

import "time"

type Associate struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	CPF             string    `json:"cpf"`
	Email           string    `json:"email"`
	Tel             string    `json:"tel"`
	DateOfBirth     time.Time `json:"date_of_birth"`    // data de nascimento
	AssociationDate time.Time `json:"association_date"` // data de associacao
	Address         string    `json:"address"`
	DonationValue   float64   `json:"donation_value"` // valor de doacao
	PaymentDate     time.Time `json:"payment_date"`   // ultima data de pagamento
	Status          bool      `json:"status"`
}
