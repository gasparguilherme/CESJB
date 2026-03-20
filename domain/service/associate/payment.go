package associate

import (
	"cesjb/domain/entities"
	"time"
)

func (s Service) CreatePayment(associateID int, date time.Time, value float64,
	status bool) (entities.Payment, error) {

	newPayment := entities.Payment{
		AssociateID: associateID,
		Date:        date,
		Value:       value,
		Status:      status,
	}

	payment, err := s.repository.SavePayment(newPayment) // PASSANDO VALOR
	if err != nil {
		return entities.Payment{}, err
	}

	return payment, nil
}
