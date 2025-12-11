package associate

import (
	"cesjb/dto/associate"
	"cesjb/handlers/associate/validate"
	"cesjb/types_"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

func (h Handler) UpdateAssociate(w http.ResponseWriter, r *http.Request) {
	// 1. Extrair ID da URL
	rawID := r.PathValue("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		slog.Error("ID inválido no path", "id", rawID, "error", err)
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	// 2. Validar ID
	if err := validate.ValidateID(id); err != nil {
		slog.Error("erro ao validar ID", "id", id, "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Decodificar body
	var input associate.UpdateAssociate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Error("erro ao decodificar JSON", "error", err)
		http.Error(w, "o JSON enviado não é valido", http.StatusBadRequest)
		return
	}

	// 5. Forçar ID do path no input para consistência
	input.ID = id

	// 6. Validar dados do DTO
	if err := validate.ValidateDTO(input); err != nil {
		slog.Error("dados inválidos", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 7. Chamar service
	updatedID, err := h.service.UpdateAssociate(input.ID, input.Name, input.Email, input.Tel, input.DateOfBirth,
		types_.DateOnly.AssociationDate, input.Address, input.DonationValue, input.PaymentDate, input.Status)
	if err != nil {
		slog.Error("erro ao atualizar associado", "error", err)
		http.Error(w, "erro ao atualizar associado", http.StatusInternalServerError)
		return
	}

	// 8. Resposta de sucesso
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Associate %d updated successfully", updatedID)))
}
