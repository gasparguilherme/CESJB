package associated

import (
	"cesjb/domain/entities"
)

func (s Service) CreateAssociated(name, cpf, email, tel string, date_of_birth, association_date string,
	address string, donation_value float64, payment_date string, status bool) (*entities.Associated, error) {

	newAssociated := entities.Associated{
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

	id, err := s.repository.SaveAssociated(newAssociated)
	if err != nil {
		return nil, err
	}

	newAssociated.ID = id
	return &newAssociated, nil
}
