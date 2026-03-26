package associate

import (
	"cesjb/domain/entities"
	"cesjb/handlers/associate/validate"
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment entities.Payment

	err := json.NewDecoder(r.Body).Decode(&payment)
	if err != nil {
		slog.Error("JSON invalido", "erro", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Ocorreu um erro inesperado",
		})
		return
	}

	err = validate.ValidatePayment(&payment)
	if err != nil {
		slog.Error("Erro ao registrar pagamento", "error", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	value, err := h.service.CreatePayment(payment.AssociateID, payment.Competence, payment.PaymentDate, payment.Value,
		payment.Status)
	if err != nil {
		slog.Error("Nao foi possivel registrar pagamento", "erro", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Nao foi possivel registrar o pagamento",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(value)

}
