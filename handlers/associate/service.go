package associate

import (
	"cesjb/domain/entities"
	"time"
)

type Service interface {
	CreateAssociate(name, cpf, email, tel string, date_of_birth, association_date time.Time,
		address string, status bool, position string) (*entities.Associate, error)

	ListAssociates() ([]entities.Associate, error)

	GetByID(id int) (*entities.Associate, error)

	UpdateAssociate(id int, name, email, tel string, dateOfBirth, associationDate time.Time,
		address string, status bool, position string) (int, error)

	GetAssociateByCPF(cpf string) (entities.Associate, error)

	FindByName(name string) ([]entities.Associate, error)

	CreatePayment(associateID int, date time.Time, value float64, status bool) (entities.Payment, error)
}
