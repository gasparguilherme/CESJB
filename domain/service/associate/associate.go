package associate

import (
	"cesjb/domain/entities"
	"cesjb/types_"
)

func (s Service) CreateAssociate(name, cpf, email, tel string, date_of_birth, association_date types_.DateOnly,
	address string, donation_value float64, payment_date types_.DateOnly, status bool) (*entities.Associate, error) {

	newAssociated := entities.Associate{
		Name:            name,
		CPF:             cpf,
		Email:           email,
		Tel:             tel,
		DateOfBirth:     date_of_birth,
		AssociationDate: association_date,
		Address:         address,
		DonationValue:   donation_value,
		PaymentDate:     payment_date,
		Status:          status,
	}

	id, err := s.repository.SaveAssociate(newAssociated)
	if err != nil {
		return nil, err
	}

	newAssociated.ID = id
	return &newAssociated, nil
}
