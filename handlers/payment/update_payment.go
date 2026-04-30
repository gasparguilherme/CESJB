package payment

import (
	"cesjb/domain"
	"cesjb/domain/entities"
	"cesjb/handlers/associate/validate"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"
)

func (h Handler) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	// 1. Extrair ID da URL
	rawID := r.PathValue("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		slog.Error("ID inválido no path", "id", rawID)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "id inválido"})
		return
	}

	// 2. Validar ID
	if err := validate.ValidateID(id); err != nil {
		slog.Error("erro ao validar ID", "id", id, "error", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// 3. Decodificar body
	var payment entities.Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		slog.Error("erro ao decodificar JSON", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "json inválido"})
		return
	}

	// 4. Forçar ID do path
	payment.ID = id

	// 5. Chamar service
	updated, err := h.service.UpdatePayment(payment.ID, payment.Competence, payment.PaymentDate, payment.Value, payment.Status)
	if err != nil {
		if errors.Is(err, domain.ErrPaymentNotFound) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}
		slog.Error("erro ao atualizar pagamento", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "erro ao atualizar pagamento"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updated)
}
