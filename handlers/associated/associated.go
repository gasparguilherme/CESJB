package associated

import (
	"cesjb/domain/entities"
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) CreateAssociated(w http.ResponseWriter, r *http.Request) {
	var associatedRequest entities.Associated

	err := json.NewDecoder(r.Body).Decode(&associatedRequest)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)
		http.Error(w, "ocorreu um erro inesperado", http.StatusInternalServerError)
		return
	}

	create, err := h.service.CreateAssociated(associatedRequest.Name, associatedRequest.CPF, associatedRequest.Email,
		associatedRequest.Tel, associatedRequest.DateOfBirth, associatedRequest.AssociationDate, associatedRequest.Address,
		associatedRequest.DonationValue, associatedRequest.PaymentDate, associatedRequest.Status)
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
