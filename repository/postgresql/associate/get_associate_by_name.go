package associate

import (
	"cesjb/domain/entities"
	"context"
	"fmt"
)

func (r Repository) FindByName(name string) ([]entities.Associate, error) {
	query := `
	SELECT id, name, cpf, email, tel, date_of_birth,
	       association_date, address,
	       status, position
	FROM associates
	WHERE name ILIKE '%' || $1 || '%'
	`

	rows, err := r.connectionInstance.Query(context.TODO(), query, name)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar associados: %w", err)
	}
	defer rows.Close()

	var associates []entities.Associate

	for rows.Next() {
		var associate entities.Associate

		err := rows.Scan(
			&associate.ID,
			&associate.Name,
			&associate.CPF,
			&associate.Email,
			&associate.Tel,
			&associate.DateOfBirth,
			&associate.AssociationDate,
			&associate.Address,
			&associate.Status,
			&associate.Position,
		)

		if err != nil {
			return nil, fmt.Errorf("erro ao ler associado: %w", err)
		}

		associates = append(associates, associate)
	}

	// 🔥 importante
	if len(associates) == 0 {
		return []entities.Associate{}, nil
	}

	return associates, nil
}
