package admin

import (
	"cesjb/authentication"
	"cesjb/domain/service/admin"
	"encoding/json"
	"log/slog"
	"net/http"
)

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginData admin.Login

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		slog.Error("não foi possivel interpretar o JSON", "error", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	userLogged, err := h.service.Login(loginData)
	if err != nil {
		slog.Error("dados de login inválido", "error", err)
		http.Error(w, "dados de login inválido", http.StatusUnauthorized)
		return
	}
	userLogged.Password = ""
	err = json.NewEncoder(w).Encode(userLogged)
	if err != nil {
		slog.Error("erro ao converter para JSON", "error", err)
		http.Error(w, "erro ao gerar resposta", http.StatusInternalServerError)
		return
	}

	token, err := authentication.CreateToken(uint64(userLogged.ID))
	if err != nil {
		slog.Error("Erro ao gerar token", "erro", err)
		http.Error(w, "Erro interno, tente novamente mais tarde", http.StatusInternalServerError)
	}
	w.Write([]byte(token))

}
