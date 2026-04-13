package associate

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) GetInactiveAssociates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	associates, err := h.service.GetInactiveAssociates()
	if err != nil {
		slog.Error("erro ao buscar associados inativos", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao buscar associados inativos",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(associates); err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
	}
}
