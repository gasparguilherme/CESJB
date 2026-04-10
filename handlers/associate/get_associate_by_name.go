package associate

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) GetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	name := r.PathValue("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "nome não informado",
		})
		return
	}

	associates, err := h.service.FindByName(name)
	if err != nil {
		slog.Error("erro ao buscar por nome", "name", name, "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao buscar associados",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(associates); err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
	}
}
