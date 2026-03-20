package associate

import (
	"time"
)

type UpdateAssociate struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Tel             string    `json:"tel"`
	DateOfBirth     time.Time `json:"date_of_birth"`    // data de nascimento
	AssociationDate time.Time `json:"association_date"` // data de associacao
	Address         string    `json:"address"`
	Status          bool      `json:"status"`
	Position        string    `json:"position"`
}
