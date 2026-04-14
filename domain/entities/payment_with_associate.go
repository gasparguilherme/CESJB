package entities

import "cesjb/types_"

// PaymentWithAssociate representa um pagamento com o nome do associado
// usado para evitar múltiplas requisições ao banco
type PaymentWithAssociate struct {
	ID            int             `json:"id"`
	AssociateID   int             `json:"associate_id"`
	AssociateName string          `json:"associate_name"`
	Competence    types_.DateOnly `json:"competence"`
	PaymentDate   types_.DateOnly `json:"payment_date"`
	Value         float64         `json:"value"`
	Status        bool            `json:"status"`
}
