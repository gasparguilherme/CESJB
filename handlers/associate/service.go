package associate

import (
	"cesjb/domain/entities"
)

type Service interface {
	CreateAssociate(name, cpf, email, tel string, date_of_birth, association_date string,
		address string, donation_value float64, payment_date string, status bool) (*entities.Associate, error)

	ListAssociates() ([]entities.Associate, error)
}
