package validate

import (
	"cesjb/dto/associate"
	"errors"
	"strings"
)

// ValidateAssociate valida os campos do DTO de atualização de associado
func ValidateDTO(input associate.UpdateAssociate) error {
	// valida campos string vazios
	switch "" {
	case strings.TrimSpace(input.Name):
		return errors.New("o nome não pode estar vazio")
	case strings.TrimSpace(input.Email):
		return errors.New("o email não pode estar vazio")
	case strings.TrimSpace(input.Tel):
		return errors.New("o telefone não pode estar vazio")
	case strings.TrimSpace(input.Address):
		return errors.New("o endereço não pode estar vazio")
	}

	// valida datas obrigatórias
	if input.DateOfBirth.IsZero() {
		return errors.New("a data de nascimento não pode estar vazia")
	}

	if input.AssociationDate.IsZero() {
		return errors.New("a data de associação não pode estar vazia")
	}

	if input.PaymentDate.IsZero() {
		return errors.New("a data de pagamento não pode estar vazia")
	}

	// valida valor da doação
	if input.DonationValue < 0 {
		return errors.New("o valor da doação não pode ser negativo")
	}

	return nil
}
