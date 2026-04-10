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

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json inválido",
		})
		return
	}

	err = validate.ValidateAdminInput(adminRequest.Name, adminRequest.Email, adminRequest.Password)
	if err != nil {
		slog.Error("erro de validação", "error", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	create, err := h.service.CreateAdmin(adminRequest.Name, adminRequest.Email, adminRequest.Password)
	if err != nil {
		slog.Error("erro ao criar administrador", "error", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao criar administrador",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(create); err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
	}

}
