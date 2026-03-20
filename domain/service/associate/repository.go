package associate

import (
	"cesjb/domain/entities"
	"cesjb/dto/associate"
)

type Repository interface {
	SaveAssociate(data entities.Associate) (int, error)
	GetAssociates() ([]entities.Associate, error)
	GetByID(id int) (*entities.Associate, error)
	UpdateAssociate(input associate.UpdateAssociate) (int, error)
	GetAssociateByCPF(cpf string) (*entities.Associate, error)
	FindByName(name string) ([]entities.Associate, error)
	SavePayment(payment entities.Payment) (entities.Payment, error)
}
