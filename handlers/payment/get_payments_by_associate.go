package payment

import (
	"cesjb/handlers/associate/validate"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

// GetPaymentsByAssociate retorna todos os pagamentos de um associado específico
// ex: GET /payments/associate/{id}
func (h Handler) GetPaymentsByAssociate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rawID := r.PathValue("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		slog.Error("erro ao converter id para inteiro", "id", rawID, "error", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id inválido",
		})
		return
	}

	if err = validate.ValidateID(id); err != nil {
		slog.Error("erro ao validar id", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id inválido",
		})
		return
	}

	payments, err := h.service.GetPaymentsByAssociate(id)
	if err != nil {
		slog.Error("erro ao buscar pagamentos do associado", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao buscar pagamentos do associado",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(payments); err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
	}
}
