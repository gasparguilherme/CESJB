package associate

import (
	"cesjb/dto/associate"
	"context"
	"fmt"
)

func (r Repository) UpdateAssociate(input associate.UpdateAssociate) (int, error) {
	query := `
        UPDATE associates
        SET 
            name = $1,
            email = $2,
            tel = $3,
            date_of_birth = $4,
            association_date = $5,
            address = $6,
            status = $7 ,
			position = $8
        WHERE id = $9  
		RETURNING id;
 
    `
	var id int
	err := r.connectionInstance.QueryRow(
		context.TODO(),
		query,
		input.Name,
		input.Email,
		input.Tel,
		input.DateOfBirth,
		input.AssociationDate,
		input.Address,
		input.Status,
		input.Position,
		input.ID,
	).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("falha em atualizar associado: %w (query: %s)", err, query)
	}
	return id, nil
}
