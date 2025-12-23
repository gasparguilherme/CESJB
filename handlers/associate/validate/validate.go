package validate

import (
	"errors"
	"strconv"
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
		return errors.New("CPF deve ter exatamente 11 dígitos")
	}

	// Verifica se todos os caracteres são números
	for _, c := range cpf {
		if c < '0' || c > '9' {
			return errors.New("CPF deve conter apenas números")
		}
	}

	// Evita CPFs com todos os dígitos iguais (ex: 11111111111)
	igual := true
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			igual = false
			break
		}
	}
	if igual {
		return errors.New("CPF inválido")
	}

	// Valida dígito verificador
	for i := 9; i <= 10; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			num, _ := strconv.Atoi(string(cpf[j]))
			sum += num * (i + 1 - j)
		}
		dig := (sum * 10) % 11
		if dig == 10 {
			dig = 0
		}
		expected, _ := strconv.Atoi(string(cpf[i]))
		if dig != expected {
			return errors.New("CPF inválido")
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
