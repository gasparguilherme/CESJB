package associated

import (
	"cesjb/domain/entities"
	"context"
	"fmt"
)

func (r Repository) SaveAssociated(data entities.Associated) (int, error) {
	query := `
    INSERT INTO associated(name, cpf, email, tel, date_of_birth, association_date, address,
    donation_value, payment_date, status)
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
    RETURNING id;
    `

	var id int
	err := r.connectionInstance.QueryRow(context.TODO(), query,
		data.Name,            // $1
		data.CPF,             // $2
		data.Email,           // $3
		data.Tel,             // $4  <-- Agora incluído
		data.DateOfBirth,     // $5
		data.AssociationDate, // $6
		data.Address,         // $7
		data.DonationValue,   // $8
		data.PaymentDate,     // $9
		data.Status,          // $10
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("executando query: %w", err)
	}
	return id, nil
}
