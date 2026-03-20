package validate

import (
	"cesjb/domain/entities"
	"errors"
	"time"
)

// Payment valida os campos do Payment
func ValidatePayment(payment *entities.Payment) error {
	if payment.AssociateID <= 0 {
		return errors.New("associateID inválido")
	}

	if payment.Value <= 0 {
		return errors.New("valor do pagamento deve ser maior que zero")
	}

	// Se a data não vier, define como agora
	if payment.Date.IsZero() {
		payment.Date = time.Now()
	}

	// Status já vem do JSON, não precisa alterar

	return nil
}
