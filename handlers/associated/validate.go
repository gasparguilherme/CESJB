package associated

import (
	"errors"
	"regexp"
	"strings"
)

func ValidateAssociate(
	name string,
	cpf string,
	email string,
	tel string,
	dateOfBirth string,
	associationDate string,
	address string,
	donationValue float64,
	paymentDate string,
) error {

	// valida campos de texto vazios
	switch "" {
	case strings.TrimSpace(name):
		return errors.New("o nome não pode estar vazio")

	case strings.TrimSpace(cpf):
		return errors.New("o cpf não pode estar vazio")

	case strings.TrimSpace(email):
		return errors.New("o email não pode estar vazio")

	case strings.TrimSpace(tel):
		return errors.New("o telefone não pode estar vazio")

	case strings.TrimSpace(dateOfBirth):
		return errors.New("a data de nascimento não pode estar vazia")

	case strings.TrimSpace(associationDate):
		return errors.New("a data de associação não pode estar vazia")

	case strings.TrimSpace(address):
		return errors.New("o endereço não pode estar vazio")

	case strings.TrimSpace(paymentDate):
		return errors.New("a data de pagamento não pode estar vazia")
	}

	// CPF: 11 dígitos
	cpfRegex := regexp.MustCompile(`^\d{11}$`)
	if !cpfRegex.MatchString(cpf) {
		return errors.New("o cpf deve conter 11 dígitos numéricos")
	}

	// email simples
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("email inválido")
	}

	// telefone somente números
	telRegex := regexp.MustCompile(`^\d{8,15}$`)
	if !telRegex.MatchString(tel) {
		return errors.New("telefone inválido: use somente números")
	}

	if donationValue < 0 {
		return errors.New("o valor da doação não pode ser negativo")
	}

	return nil
}
