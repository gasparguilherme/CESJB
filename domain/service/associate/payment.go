package associate

import (
	"cesjb/domain"
	"cesjb/domain/entities"
	"cesjb/types_"
)

func (s Service) CreatePayment(associateID int, competence, paymentDate types_.DateOnly, value float64,
	status bool) (entities.Payment, error) {

	if associateID <= 0 {
		return entities.Payment{}, domain.ErrAssociateNotFound
	}

	if value <= 0 {
		return entities.Payment{}, domain.ErrInvalidValue
	}
	newPayment := entities.Payment{
		AssociateID: associateID,
		PaymentDate: paymentDate,
		Value:       value,
		Status:      status,
		Competence:  competence,
	}

	payment, err := s.repository.SavePayment(newPayment)
	if err != nil {
		return entities.Payment{}, err
	}

	return payment, nil
}
