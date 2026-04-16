package associate

import (
	"cesjb/domain"
	"cesjb/domain/entities"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r Repository) SaveAssociate(data entities.Associate) (int, error) {
	query := `
    INSERT INTO associates(name, cpf, email, tel, date_of_birth, association_date, address,
    status, position)
    VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
    RETURNING id;
    `

	var id int
	err := r.connectionInstance.QueryRow(context.TODO(), query,
		data.Name,
		data.CPF,
		data.Email,
		data.Tel,
		time.Time(data.DateOfBirth),
		time.Time(data.AssociationDate),
		data.Address,
		data.Status,
		data.Position,
	).Scan(&id)

	if err != nil {
		// verifica se é erro de CPF duplicado (unique constraint)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, domain.ErrCPFAlreadyExists
		}

		return 0, fmt.Errorf("executando query: %w", err)
	}

	return id, nil
}
