package associate

import (
	"cesjb/domain/entities"
	"fmt"
)

func (s Service) GetAssociateByCPF(cpf string) (entities.Associate, error) {
	associate, err := s.repository.GetAssociateByCPF(cpf)
	if err != nil {
		return entities.Associate{}, fmt.Errorf("erro ao buscar associado: %w", err)
	}

	return *associate, nil
}
