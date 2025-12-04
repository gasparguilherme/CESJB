package associate

import (
	"cesjb/domain/entities"
	"time"
)

type Service interface {
	CreateAssociate(name, cpf, email, tel string, date_of_birth, association_date time.Time,
		address string, donation_value float64, payment_date time.Time, status bool) (*entities.Associate, error)

	ListAssociates() ([]entities.Associate, error)
}
