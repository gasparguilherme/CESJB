package associate

import (
	"cesjb/handlers/associate/validate"
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) GetAssociates(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	cpf := r.URL.Query().Get("cpf")

	w.Header().Set("Content-Type", "application/json")

	// filtro por CPF
	if cpf != "" {
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
		json.NewEncoder(w).Encode(associate)
		return
	}

	// filtro por nome
	if name != "" {
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
		json.NewEncoder(w).Encode(associates)
		return
	}

	// sem filtro
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
	json.NewEncoder(w).Encode(associates)
}
