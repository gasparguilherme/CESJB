package validate

import (
	"errors"
	"strconv"
)

// ValidarCPF verifica se o CPF é válido
func ValidateCPF(cpf string) error {
	// Remove caracteres não numéricos (caso queira aceitar com pontos e traço)
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

	return nil
}
