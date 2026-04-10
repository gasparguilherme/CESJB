package associate

import (
	"cesjb/domain/entities"
	"cesjb/handlers/associate/validate"
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) CreateAssociate(w http.ResponseWriter, r *http.Request) {
	var associatedRequest entities.Associate

	err := json.NewDecoder(r.Body).Decode(&associatedRequest)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Ocorreu um erro inesperado",
		})
		return
	}

	// VALIDAÇÃO: verifique o erro aqui!
	err = validate.ValidateAssociate(associatedRequest.Name, associatedRequest.CPF, associatedRequest.Email,
		associatedRequest.Tel, associatedRequest.DateOfBirth, associatedRequest.AssociationDate, associatedRequest.Address,
		associatedRequest.Position)

	if err != nil {
		slog.Error("erro de validação", "error", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	create, err := h.service.CreateAssociate(associatedRequest.Name, associatedRequest.CPF, associatedRequest.Email,
		associatedRequest.Tel, associatedRequest.DateOfBirth, associatedRequest.AssociationDate, associatedRequest.Address,
		associatedRequest.Status, associatedRequest.Position)
	if err != nil {
		slog.Error("erro ao criar usuario", "error", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao criar associado",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(create); err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
	}
}
