package payment

import (
	"cesjb/domain/entities"
	"fmt"
)

// GetPaymentsByAssociate retorna todos os pagamentos de um associado específico
func (s Service) GetPaymentsByAssociate(associateID int) ([]entities.Payment, error) {
	payments, err := s.repository.GetPaymentsByAssociate(associateID)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar pagamentos do associado: %w", err)
	}

	return payments, nil
}
