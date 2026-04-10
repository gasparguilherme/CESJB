package payment

import (
	"cesjb/domain"
	"cesjb/domain/entities"
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return entities.Payment{}, domain.ErrPaymentAlreadyExists
		}

		return entities.Payment{}, domain.ErrCreatePayment
	}

	return payment, nil
}
