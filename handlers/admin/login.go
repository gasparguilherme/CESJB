package admin

import (
	"cesjb/authentication"
	"cesjb/domain/service/admin"
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginData admin.Login

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "json inválido",
		})
		return
	}

	userLogged, err := h.service.Login(loginData)
	if err != nil {
		slog.Error("dados de login inválido", "error", err)

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "dados de login inválidos",
		})
		return
	}

	userLogged.Password = ""

	token, err := authentication.CreateToken(uint64(userLogged.ID))
	if err != nil {
		slog.Error("erro ao gerar token", "error", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "erro interno ao gerar token",
		})
		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]any{
		"user":  userLogged,
		"token": token,
	})
}
