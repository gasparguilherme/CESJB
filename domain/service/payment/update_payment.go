package payment

import (
	"cesjb/domain/entities"
	"cesjb/types_"
)

func (s Service) UpdatePayment(id int, competence, paymentDate types_.DateOnly, value float64,
	status bool) (entities.Payment, error) {

	updated := entities.Payment{
		ID:          id,
		Competence:  competence,
		PaymentDate: paymentDate,
		Value:       value,
		Status:      status,
	}

	payment, err := s.repository.UpdatePayment(updated)
	if err != nil {
		return entities.Payment{}, err
	}

	return payment, nil
}
