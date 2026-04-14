package payment

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

// GetPaymentsByAssociate retorna todos os pagamentos de um associado específico
// ex: GET /payments/associate/{id}
func (h Handler) GetPaymentsByAssociate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")
	if idStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id não informado",
		})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
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
