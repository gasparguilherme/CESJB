package payment

import (
	"cesjb/types_"
	"context"
	"fmt"
	"time"
)

// PaymentWithAssociate representa um pagamento com o nome do associado
// usado para evitar múltiplas requisições ao banco
type PaymentWithAssociate struct {
	ID            int
	AssociateID   int
	AssociateName string
	Competence    types_.DateOnly
	PaymentDate   types_.DateOnly
	Value         float64
	Status        bool
}

// GetPaymentsByMonth retorna todos os pagamentos de um mês específico
// junto com o nome do associado via JOIN
func (r Repository) GetPaymentsByMonth(competence types_.DateOnly) ([]PaymentWithAssociate, error) {
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
	payments := []PaymentWithAssociate{}

	for rows.Next() {
		var p PaymentWithAssociate
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
