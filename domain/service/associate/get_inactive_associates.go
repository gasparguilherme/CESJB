package associate

import (
	"cesjb/domain/entities"
	"fmt"
)

func (s Service) GetInactiveAssociates() ([]entities.Associate, error) {
	associates, err := s.repository.GetInactiveAssociates()
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar associados inativos: %w", err)
	}

	return associates, nil
}
