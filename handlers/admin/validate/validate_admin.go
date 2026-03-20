package validate

import (
	"errors"
	"net/mail"
	"strings"
)

func ValidateAdminInput(name, email, password string) error {
	// Nome não pode ser vazio
	if strings.TrimSpace(name) == "" {
		return errors.New("nome é obrigatório")
	}

	// Email não pode ser vazio e deve ter formato válido
	if strings.TrimSpace(email) == "" {
		return errors.New("email é obrigatório")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("email inválido")
	}

	// Senha não pode ser vazia e deve ter pelo menos 6 caracteres
	if len(password) < 6 {
		return errors.New("senha deve ter pelo menos 6 caracteres")
	}

	return nil
}
