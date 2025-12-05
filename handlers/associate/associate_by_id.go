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
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}
	err = validate.ValidateID(id)
	if err != nil {
		slog.Error("erro ao buscar ID", "error", err)
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return

	}

	associate, err := h.service.GetByID(id)
	if err != nil {
		slog.Error("erro ao buscar associado", "error", err)
		http.Error(w, "erro interno, falha ao buscar associado", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(associate)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		http.Error(w, "erro interno", http.StatusInternalServerError)
		return
	}

}
