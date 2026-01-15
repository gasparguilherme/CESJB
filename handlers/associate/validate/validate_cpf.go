package validate

import (
	"errors"
	"strconv"
)

// ValidateCPF valida se o CPF é válido (somente números, 11 dígitos e DV)
func ValidateCPF(cpf string) error {
	if len(cpf) != 11 {
		return errors.New("CPF deve ter exatamente 11 dígitos")
	}

	for _, c := range cpf {
		if c < '0' || c > '9' {
			return errors.New("CPF deve conter apenas números")
		}
	}

	// Evita CPFs com todos os dígitos iguais
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

	// Validação dos dígitos verificadores
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

	return nil
}
