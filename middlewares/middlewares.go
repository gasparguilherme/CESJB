package middlewares

import (
	"cesjb/authentication"
	"log"
	"log/slog"
	"net/http"
)

// logger imprime no terminal informacões da requisição no terminal
func Logger(nexFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nexFunc(w, r)
	}
}

// authenticate verifica se o usuario fazendo requisição está autenticado
func Authenticate(nexFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			slog.Error("Erro ao autenticar token", "error", err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return

		}
		nexFunc(w, r)
	}
}
