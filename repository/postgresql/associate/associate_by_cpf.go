package associate

import (
	"cesjb/domain/entities"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func (r Repository) GetAssociateByCPF(cpf string) (*entities.Associate, error) {
	query := `
        SELECT id, name, cpf, email, tel, date_of_birth, association_date,
               address, donation_value, payment_date, status
        FROM associates
        WHERE cpf = $1
    `

	row := r.connectionInstance.QueryRow(context.TODO(), query, cpf)

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
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("associado não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar associado: %w", err)
	}

	return &associate, nil
}
