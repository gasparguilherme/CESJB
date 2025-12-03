package entities

type Associated struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	CPF             string  `json:"cpf"`
	Email           string  `json:"email"`
	Tel             string  `json:"tel"`
	DateOfBirth     string  `json:"date_of_birth"`    // data de nascimento
	AssociationDate string  `json:"association_date"` // data de associacao
	Address         string  `json:"address"`
	DonationValue   float64 `json:"donation_value"` // valor da doacao
	PaymentDate     string  `json:"payment_date"`   // ultima data de pagamento
	Status          bool    `json:"status"`
}
