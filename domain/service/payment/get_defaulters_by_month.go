package payment

import (
	"cesjb/domain/entities"
	"cesjb/types_"
	"fmt"
)

// GetDefaultersByMonth retorna os associados ativos que não pagaram no mês informado
func (s Service) GetDefaultersByMonth(competence types_.DateOnly) ([]entities.Associate, error) {
	defaulters, err := s.repository.GetDefaultersByMonth(competence)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar inadimplentes do mês: %w", err)
	}

	return defaulters, nil
}
