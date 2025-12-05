package validate

import (
	"errors"
)

func ValidateID(id int) error {
	if id <= 0 {
		return errors.New("ID inválido: deve ser maior que zero")
	}
	return nil
}
