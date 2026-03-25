package entities

import (
	"cesjb/types_"
)

type Associate struct {
	ID              int             `json:"id"`
	Name            string          `json:"name"`
	CPF             string          `json:"cpf"`
	Email           string          `json:"email"`
	Tel             string          `json:"tel"`
	DateOfBirth     types_.DateOnly `json:"date_of_birth"`    // data de nascimento
	AssociationDate types_.DateOnly `json:"association_date"` // data de associacao
	Address         string          `json:"address"`
	Status          bool            `json:"status"`
	Position        string          `json:"position"` //cargo
}
