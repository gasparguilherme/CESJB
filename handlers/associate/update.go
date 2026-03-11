package associate

import (
	"cesjb/dto/associate"
	"cesjb/handlers/associate/validate"
	"encoding/json"
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

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "id inválido",
		})
		return
	}

	// 2. Validar ID
	if err := validate.ValidateID(id); err != nil {
		slog.Error("erro ao validar ID", "id", id, "error", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	// 3. Decodificar body
	var input associate.UpdateAssociate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		slog.Error("erro ao decodificar JSON", "error", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json inválido",
		})
		return
	}

	// 5. Forçar ID do path no input para consistência
	input.ID = id

	// 6. Validar dados do DTO
	if err := validate.ValidateDTO(input); err != nil {
		slog.Error("dados inválidos", "error", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	// 7. Chamar service
	updatedID, err := h.service.UpdateAssociate(input.ID, input.Name, input.Email, input.Tel, input.DateOfBirth,
		input.AssociationDate, input.Address, input.DonationValue, input.PaymentDate, input.Status, input.Position)
	if err != nil {
		slog.Error("erro ao atualizar associado", "error", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro ao atualizar associado",
		})
		return
	}

	// 8. Resposta de sucesso
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "associado atualizado com sucesso",
		"id":      updatedID,
	})
}
