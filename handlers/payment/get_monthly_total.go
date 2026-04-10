package payment

import (
	"cesjb/types_"
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) GetMonthlyTotal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// pega o query param competence ex: ?competence=2026-04-01
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

	total, err := h.service.GetMonthlyTotal(competence)
	if err != nil {
		slog.Error("erro ao buscar total mensal", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao buscar total de pagamentos",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(map[string]float64{
		"total": total,
	}); err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
	}
}
