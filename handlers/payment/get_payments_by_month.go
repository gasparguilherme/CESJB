package payment

import (
	"cesjb/types_"
	"encoding/json"
	"log/slog"
	"net/http"
)

// GetPaymentsByMonth retorna os pagamentos de um mês específico
// recebe o query param competence no formato yyyy-mm-dd
// ex: GET /payments?competence=2026-04-01
func (h Handler) GetPaymentsByMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// pega o query param competence
	competenceStr := r.URL.Query().Get("competence")
	if competenceStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "competência não informada",
		})
		return
	}

	// converte a string para DateOnly
	var competence types_.DateOnly
	if err := competence.UnmarshalJSON([]byte(`"` + competenceStr + `"`)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "formato de competência inválido, use yyyy-mm-dd",
		})
		return
	}

	payments, err := h.service.GetPaymentsByMonth(competence)
	if err != nil {
		slog.Error("erro ao buscar pagamentos do mês", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao buscar pagamentos",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(payments); err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
	}
}
