package associated

import (
	"cesjb/domain/entities"
)

type Service interface {
	CreateAssociated(name, cpf, email, tel string, date_of_birth, association_date string,
		address string, donation_value float64, payment_date string, status bool) (*entities.Associated, error)
}
