package associate

import (
	"cesjb/domain/entities"
	"cesjb/types_"
	"time"
)

type Service interface {
	CreateAssociate(name, cpf, email, tel string, date_of_birth, association_date time.Time,
		address string, donation_value float64, payment_date time.Time, status bool,
		position string) (*entities.Associate, error)

	ListAssociates() ([]entities.Associate, error)

	GetByID(id int) (*entities.Associate, error)

	UpdateAssociate(id int, name, email, tel string, dateOfBirth, associationDate types_.DateOnly,
		address string, donationValue float64, paymentDate types_.DateOnly, status bool) (int, error)

	GetAssociateByCPF(cpf string) (entities.Associate, error)
}
