package validate

import (
	"errors"
	"strings"
	"time"
)

func ValidateAssociate(
	name string,
	cpf string,
	email string,
	tel string,
	dateOfBirth time.Time,
	associationDate time.Time,
	address string,
	donationValue float64,
	paymentDate time.Time,
	position string,
) error {

	// valida campos string vazios
	switch "" {
	case strings.TrimSpace(name):
		return errors.New("o nome não pode estar vazio")

	case strings.TrimSpace(cpf):
		return errors.New("o cpf não pode estar vazio")

	case strings.TrimSpace(email):
		return errors.New("o email não pode estar vazio")

	case strings.TrimSpace(tel):
		return errors.New("o telefone não pode estar vazio")

	case strings.TrimSpace(address):
		return errors.New("o endereço não pode estar vazio")

	case strings.TrimSpace(position):
		return errors.New("o cargo não pode estar vazio")
	}

	// valida cpf sem pontos
	if len(cpf) != 11 {
		return errors.New("o cpf deve conter exatamente 11 dígitos")
	}

	for _, r := range cpf {
		if r < '0' || r > '9' {
			return errors.New("o cpf deve conter apenas números")
		}
	}

	// valida datas obrigatórias
	if dateOfBirth.IsZero() {
		return errors.New("a data de nascimento não pode estar vazia")
	}

	if associationDate.IsZero() {
		return errors.New("a data de associação não pode estar vazia")
	}

	if paymentDate.IsZero() {
		return errors.New("a data de pagamento não pode estar vazia")
	}

	return nil
}
