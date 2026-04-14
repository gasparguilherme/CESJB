package payment

import (
	"cesjb/domain/entities"
	"cesjb/types_"
	"fmt"
)

// GetPaymentsByMonth retorna os pagamentos de um mês específico
// com o nome do associado incluído
func (s Service) GetPaymentsByMonth(competence types_.DateOnly) ([]entities.PaymentWithAssociate, error) {
	payments, err := s.repository.GetPaymentsByMonth(competence)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar pagamentos do mês: %w", err)
	}

	return payments, nil
}
