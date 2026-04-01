package payment

import (
	"cesjb/domain"
	"cesjb/domain/entities"
	"cesjb/handlers/payment/validate"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

func (h Handler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment entities.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "JSON inválido",
		})
		return
	}

	err = validate.ValidatePayment(&payment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	value, err := h.service.CreatePayment(
		payment.AssociateID,
		payment.Competence,
		payment.PaymentDate,
		payment.Value,
		payment.Status,
	)

	if err != nil {

		if errors.Is(err, domain.ErrPaymentAlreadyExists) {
			slog.Error("Erro ao registrar pagamento", "error", err)
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Erro interno ao registrar pagamento",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(value)
}
