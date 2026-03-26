package domain

import "errors"

var (
	ErrPaymentAlreadyExists = errors.New("pagamento já registrado para essa competência")
	ErrAssociateNotFound    = errors.New("associado não encontrado")
	ErrInvalidValue         = errors.New("valor de pagamento inválido")
	ErrCreatePayment        = errors.New("erro ao registrar pagamento")
)
