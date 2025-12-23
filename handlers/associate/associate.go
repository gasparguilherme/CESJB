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
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}

	// VALIDAÇÃO: verifique o erro aqui!
	err = validate.ValidateAssociate(associatedRequest.Name, associatedRequest.CPF, associatedRequest.Email,
		associatedRequest.Tel, associatedRequest.DateOfBirth, associatedRequest.AssociationDate, associatedRequest.Address,
		associatedRequest.DonationValue, associatedRequest.PaymentDate, associatedRequest.Position)

	if err != nil {
		slog.Error("erro de validação", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest) // ← Retorne o erro de validação
		return
	}

	create, err := h.service.CreateAssociate(associatedRequest.Name, associatedRequest.CPF, associatedRequest.Email,
		associatedRequest.Tel, associatedRequest.DateOfBirth, associatedRequest.AssociationDate, associatedRequest.Address,
		associatedRequest.DonationValue, associatedRequest.PaymentDate, associatedRequest.Status, associatedRequest.Position)
	if err != nil {
		slog.Error("erro ao criar usuario", "error", err)
		http.Error(w, "erro ao criar usuario", http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(create)
	if err != nil {
		slog.Error("erro ao converter para formato JSON", "error", err)
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}
}
