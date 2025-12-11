package admin

import (
	"cesjb/domain/entities"
	"context"
	"fmt"
)

func (r Repository) SaveAdmin(data entities.Admin) (int, error) {
	query := `
	INSERT INTO admin(name, email, password)
    VALUES($1, $2, $3)
    RETURNING id;
	`
	var id int

	err := r.connectionInstance.QueryRow(context.TODO(), query,
		data.Name,
		data.Email,
		data.Password).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("executando query: %w", err)

	}
	return id, nil
}
