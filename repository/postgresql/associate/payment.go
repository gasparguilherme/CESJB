package associate

import (
	"cesjb/domain/entities"
	"context"
	"fmt"
	"strings"
	"time"
)

func (r Repository) SavePayment(payment entities.Payment) (entities.Payment, error) {

	query := `
	INSERT INTO payments (
		associate_id,
		competence,
		payment_date,
		value,
		status
	)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	err := r.connectionInstance.QueryRow(
		context.TODO(),
		query,
		payment.AssociateID,
		time.Time(payment.Competence),
		time.Time(payment.PaymentDate),
		payment.Value,
		payment.Status,
	).Scan(&payment.ID)

	if err != nil {

		// erro de pagamento duplicado
		if strings.Contains(err.Error(), "unique_associate_competence") {
			return entities.Payment{}, fmt.Errorf("já existe um pagamento registrado para este associado neste mês")
		}

		// erro genérico
		return entities.Payment{}, fmt.Errorf("erro ao registrar pagamento")
	}

	return payment, nil
}
