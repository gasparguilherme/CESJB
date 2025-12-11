package admin

import (
	"cesjb/domain/entities"
	"cesjb/handlers/admin/validate"
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) CreateAdmin(w http.ResponseWriter, r *http.Request) {
	var adminRequest entities.Admin

	err := json.NewDecoder(r.Body).Decode(&adminRequest)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}

	err = validate.ValidateAdminInput(adminRequest.Name, adminRequest.Email, adminRequest.Password)
	if err != nil {
		slog.Error("erro de validação", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	create, err := h.service.CreateAdmin(adminRequest.Name, adminRequest.Email, adminRequest.Password)
	if err != nil {
		slog.Error("erro ao criar administrador", "error", err)
		http.Error(w, "erro ao criar administrador", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(create)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}

}
