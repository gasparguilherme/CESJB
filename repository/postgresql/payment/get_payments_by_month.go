package payment

import (
	"cesjb/domain/entities"
	"cesjb/types_"
	"context"
	"fmt"
	"time"
)

// GetPaymentsByMonth retorna todos os pagamentos de um mês específico
// junto com o nome do associado via JOIN
func (r Repository) GetPaymentsByMonth(competence types_.DateOnly) ([]entities.PaymentWithAssociate, error) {
	query := `
		SELECT 
			p.id,
			p.associate_id,
			a.name,
			p.competence,
			p.payment_date,
			p.value,
			p.status
		FROM payments p
		INNER JOIN associates a ON a.id = p.associate_id
		WHERE DATE_TRUNC('month', p.competence::date) = DATE_TRUNC('month', $1::date)
		ORDER BY a.name ASC
	`

	rows, err := r.connectionInstance.Query(context.TODO(), query, time.Time(competence))
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar pagamentos: %w", err)
	}
	defer rows.Close()

	// inicializa slice vazio para evitar retorno null quando não há pagamentos
	payments := []entities.PaymentWithAssociate{}

	for rows.Next() {
		var p entities.PaymentWithAssociate
		err := rows.Scan(
			&p.ID,
			&p.AssociateID,
			&p.AssociateName,
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
