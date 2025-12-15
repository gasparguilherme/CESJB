package admin

import (
	"cesjb/domain/entities"
	"context"
	"fmt"
)

func (r Repository) FindUserByEmail(email string) (*entities.Admin, error) {

	var admin entities.Admin

	query := `
	SELECT id, name, email, password FROM admin WHERE email = $1
	`
	err := r.connectionInstance.QueryRow(context.TODO(), query, email).Scan(
		&admin.ID,
		&admin.Name,
		&admin.Email,
		&admin.Password,
	)

	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado: %w", err)
	}

	return &admin, nil
}
