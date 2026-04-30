package payment

import (
	"cesjb/domain"
	"cesjb/domain/entities"
	"context"
	"time"
)

func (r Repository) UpdatePayment(payment entities.Payment) (entities.Payment, error) {
	query := `
	UPDATE payments SET
		competence   = $1,
		payment_date = $2,
		value        = $3,
		status       = $4
	WHERE id = $5
	RETURNING id, associate_id, competence, payment_date, value, status
	`
	err := r.connectionInstance.QueryRow(
		context.TODO(),
		query,
		time.Time(payment.Competence),
		time.Time(payment.PaymentDate),
		payment.Value,
		payment.Status,
		payment.ID,
	).Scan(
		&payment.ID,
		&payment.AssociateID,
		&payment.Competence,
		&payment.PaymentDate,
		&payment.Value,
		&payment.Status,
	)
	if err != nil {
		return entities.Payment{}, domain.ErrPaymentNotFound
	}

	return payment, nil
}
