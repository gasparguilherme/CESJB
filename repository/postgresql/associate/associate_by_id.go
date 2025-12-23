package associate

import (
	"cesjb/domain/entities"
	"context"
	"fmt"
)

func (r Repository) GetByID(id int) (*entities.Associate, error) {
	query := `
        SELECT id, name, cpf, email, tel, date_of_birth, association_date,
               address, donation_value, payment_date, status, position
        FROM associates 
        WHERE id = $1;
    `

	row := r.connectionInstance.QueryRow(context.TODO(), query, id)

	var associate entities.Associate

	err := row.Scan(
		&associate.ID,
		&associate.Name,
		&associate.CPF,
		&associate.Email,
		&associate.Tel,
		&associate.DateOfBirth,
		&associate.AssociationDate,
		&associate.Address,
		&associate.DonationValue,
		&associate.PaymentDate,
		&associate.Status,
		&associate.Position,
	)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar associado: %w", err)
	}

	return &associate, nil
}
