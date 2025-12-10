package associate

import (
	"cesjb/handlers/associate/validate"
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
)

func (h Handler) GetAssociateByCPF(w http.ResponseWriter, r *http.Request) {
	cpf := strings.TrimPrefix(r.URL.Path, "/associate/cpf/")

	err := validate.ValidateCPF(cpf)
	if err != nil {
		slog.Error("erro ao validar CPF", "cpf", cpf, "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	associate, err := h.service.GetAssociateByCPF(cpf)
	if err != nil {
		http.Error(w, "Associado não encontrado", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(associate); err != nil {
		slog.Error("Erro ao gerar JSON", "error", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return

	}
}
