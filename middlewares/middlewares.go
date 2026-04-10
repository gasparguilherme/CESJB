package middlewares

import (
	"cesjb/authentication"
	"log"
	"log/slog"
	"net/http"
)

// Logger imprime no terminal informações da requisição
func Logger(nexFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nexFunc(w, r)
	}
}

// Authenticate verifica se o usuário fazendo requisição está autenticado
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

// CORS libera o acesso do browser à API
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// requisições OPTIONS são enviadas pelo browser antes da requisição real (preflight)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
