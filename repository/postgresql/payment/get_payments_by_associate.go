package payment

import (
	"cesjb/domain/entities"
	"context"
	"fmt"
)

// GetPaymentsByAssociate retorna todos os pagamentos de um associado específico
func (r Repository) GetPaymentsByAssociate(associateID int) ([]entities.Payment, error) {
	query := `
			SELECT
			id,
			associate_id,
			competence,
			payment_date,
			value,
			status
		FROM payments
		WHERE associate_id = $1
		AND competence >= DATE_TRUNC('month', NOW()) - INTERVAL '11 months'
		ORDER BY competence DESC
	`

	rows, err := r.connectionInstance.Query(context.TODO(), query, associateID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar pagamentos do associado: %w", err)
	}
	defer rows.Close()

	payments := []entities.Payment{}

	for rows.Next() {
		var p entities.Payment
		err := rows.Scan(
			&p.ID,
			&p.AssociateID,
			&p.Competence,
			&p.PaymentDate,
			&p.Value,
			&p.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler pagamento: %w", err)
		}
		payments = append(payments, p)
	}

	return payments, nil
}
