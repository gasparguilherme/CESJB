package entities

import (
	"cesjb/types_"
)

type Payment struct {
	ID          int
	AssociateID int             `json:"associateID"`
	Date        types_.DateOnly `json:"date"`
	Value       float64         `json:"value"`
	Status      bool            `json:"status"`
}
