package associate

import (
	"cesjb/handlers/associate/validate"
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) GetByCPF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cpf := r.PathValue("cpf")

	if err := validate.ValidateCPF(cpf); err != nil {
		slog.Error("erro ao validar CPF", "cpf", cpf, "error", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	associate, err := h.service.GetAssociateByCPF(cpf)
	if err != nil {
		slog.Error("associado não encontrado", "cpf", cpf, "error", err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "associado não encontrado",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(associate); err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
	}
}
