package payment

import (
	"cesjb/types_"
	"fmt"
)

func (s Service) GetMonthlyTotal(competence types_.DateOnly) (float64, error) {
	total, err := s.repository.GetMonthlyTotal(competence)
	if err != nil {
		return 0, fmt.Errorf("erro ao buscar total mensal: %w", err)
	}

	return total, nil
}
