package entities

import "time"

type Payment struct {
	ID          int
	AssociateID int       `json:"associateID"`
	Date        time.Time `json:"date"`
	Value       float64   `json:"value"`
	Status      bool      `json:"status"`
}
