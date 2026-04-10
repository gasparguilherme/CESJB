package associate

import (
	"cesjb/handlers/associate/validate"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func (h Handler) GetByID(w http.ResponseWriter, r *http.Request) {
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
	err = validate.ValidateID(id)
	if err != nil {
		slog.Error("erro ao buscar ID", "error", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id inválido",
		})
		return

	}

	associate, err := h.service.GetByID(id)
	if err != nil {
		slog.Error("erro ao buscar associado", "error", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao buscar associado",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(associate); err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
	}
}
