package associate

import (
	"cesjb/domain/entities"
	"context"
	"fmt"
)

func (r Repository) SaveAssociate(data entities.Associate) (int, error) {
	query := `
    INSERT INTO associates(name, cpf, email, tel, date_of_birth, association_date, address,
    donation_value, payment_date, status)
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    RETURNING id;
    `

	var id int
	err := r.connectionInstance.QueryRow(context.TODO(), query,
		data.Name,
		data.CPF,
		data.Email,
		data.Tel,
		data.DateOfBirth,
		data.AssociationDate,
		data.Address,
		data.DonationValue,
		data.PaymentDate,
		data.Status,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("executando query: %w", err)
	}
	return id, nil
}
