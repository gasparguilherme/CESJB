package associate

import (
	"cesjb/domain/entities"
	"errors"
)

func (s Service) FindByName(name string) ([]entities.Associate, error) {
	if name == "" {
		return nil, errors.New("nome nao pode estar vazio")
	}
	return s.repository.FindByName(name)
}
