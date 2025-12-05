package associate

import (
	"cesjb/dto"
	"errors"
	"fmt"
	"time"
)

func (s Service) UpdateAssociate(id int, name, email, tel string, dateOfBirth, associationDate time.Time,
	address string, donationValue float64, paymentDate time.Time, status bool) (int, error) {
	associate := dto.UpdateAssociate{
		ID:              id,
		Name:            name,
		Email:           email,
		Tel:             tel,
		DateOfBirth:     dateOfBirth,
		AssociationDate: associationDate,
		Address:         address,
		DonationValue:   donationValue,
		PaymentDate:     paymentDate,
		Status:          status,
	}
	updateID, err := s.repository.UpdateAssociate(associate)
	if err != nil {
		return 0, fmt.Errorf("falha em atualizar associado: %w", err)
	}
	if updateID != id {
		return 0, errors.New("inconsistency in update: ID modified")
	}
	return updateID, nil

}
