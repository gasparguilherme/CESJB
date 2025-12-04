package associate

import (
	"encoding/json"
	"net/http"
)

func (h Handler) GetAssociates(w http.ResponseWriter, r *http.Request) {
	associates, err := h.service.ListAssociates()
	if err != nil {
		http.Error(w, "Erro ao buscar associados: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(associates); err != nil {
		http.Error(w, "Erro ao gerar resposta JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
