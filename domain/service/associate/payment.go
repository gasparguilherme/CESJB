package associate

import (
	"cesjb/domain/entities"
	"cesjb/types_"
)

func (s Service) CreatePayment(associateID int, competence, date types_.DateOnly, value float64,
	status bool) (entities.Payment, error) {

	newPayment := entities.Payment{
		AssociateID: associateID,
		Date:        date,
		Value:       value,
		Status:      status,
		Competence:  competence,
	}

	payment, err := s.repository.SavePayment(newPayment) // PASSANDO VALOR
	if err != nil {
		return entities.Payment{}, err
	}

	return payment, nil
}
