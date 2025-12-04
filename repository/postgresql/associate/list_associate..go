package associate

import (
	"cesjb/domain/entities"
	"context"
	"fmt"
)

func (r Repository) GetAssociates() ([]entities.Associate, error) {
	query := `
        SELECT 
            id,
            name,
            cpf,
            email,
            tel,
            date_of_birth,
            association_date,
            address,
            donation_value,
            payment_date,
            status
        FROM associates
        ORDER BY name ASC;
    `

	rows, err := r.connectionInstance.Query(context.TODO(), query)
	if err != nil {
		return nil, fmt.Errorf("executando query: %w", err)
	}
	defer rows.Close()

	var associates []entities.Associate

	for rows.Next() {
		var a entities.Associate
		err := rows.Scan(
			&a.ID,
			&a.Name,
			&a.CPF,
			&a.Email,
			&a.Tel,
			&a.DateOfBirth,
			&a.AssociationDate,
			&a.Address,
			&a.DonationValue,
			&a.PaymentDate,
			&a.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("scanning row: %w", err)
		}
		associates = append(associates, a)
	}

	return associates, nil
}
