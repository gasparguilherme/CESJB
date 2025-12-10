package associate

import (
	"cesjb/domain/entities"
	"cesjb/dto"
)

type Repository interface {
	SaveAssociate(data entities.Associate) (int, error)
	GetAssociates() ([]entities.Associate, error)
	GetByID(id int) (*entities.Associate, error)
	UpdateAssociate(input dto.UpdateAssociate) (int, error)
	GetAssociateByCPF(cpf string) (*entities.Associate, error)
}
