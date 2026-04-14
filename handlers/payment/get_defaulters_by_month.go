package payment

import (
	"cesjb/types_"
	"encoding/json"
	"log/slog"
	"net/http"
)

// GetDefaultersByMonth retorna os associados ativos que não pagaram no mês informado
// recebe o query param competence no formato yyyy-mm-dd
// ex: GET /payments/defaulters?competence=2026-04-01
func (h Handler) GetDefaultersByMonth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	competenceStr := r.URL.Query().Get("competence")
	if competenceStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "competência não informada",
		})
		return
	}

	var competence types_.DateOnly
	if err := competence.UnmarshalJSON([]byte(`"` + competenceStr + `"`)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "formato de competência inválido, use yyyy-mm-dd",
		})
		return
	}

	defaulters, err := h.service.GetDefaultersByMonth(competence)
	if err != nil {
		slog.Error("erro ao buscar inadimplentes do mês", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao buscar inadimplentes",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(defaulters); err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
	}
}
