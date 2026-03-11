package associate

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) GetAssociates(w http.ResponseWriter, r *http.Request) {
	associates, err := h.service.ListAssociates()
	if err != nil {
		slog.Error("erro ao buscar associados", "error", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao buscar associados",
		})
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(associates); err != nil {
		slog.Error("erro ao gerar resposta JSON", "error", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao gerar resposta JSON",
		})
		return
	}

}
