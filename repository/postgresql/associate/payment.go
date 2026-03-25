package associate

import (
	"cesjb/domain/entities"
	"context"
	"fmt"
	"time"
)

func (r Repository) SavePayment(payment entities.Payment) (entities.Payment, error) {
	query := `
        INSERT INTO payments (associate_id, date, value, status)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
	err := r.connectionInstance.QueryRow(context.TODO(), query,
		payment.AssociateID,
		time.Time(payment.Date),
		payment.Value,
		payment.Status,
	).Scan(&payment.ID)
	if err != nil {
		return entities.Payment{}, fmt.Errorf("executando query %w", err)
	}
	return payment, nil
}
