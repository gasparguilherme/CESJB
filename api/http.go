package api

import (
	"cesjb/config"
	"cesjb/middlewares"
	"log/slog"
	"net/http"
)

func StartApp(associateHandler Associate, adminHandler Admin, paymentHandler Payment) {
	mux := http.NewServeMux()

	// Rotas Associado
	mux.Handle("POST /associate", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(associateHandler.CreateAssociate))))

	mux.Handle("GET /associates", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(associateHandler.GetAssociates))))

	mux.Handle("GET /associates/inactive", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(associateHandler.GetInactiveAssociates))))

	mux.Handle("GET /associate/id/{id}", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(associateHandler.GetByID))))

	mux.Handle("GET /associate/cpf/{cpf}", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(associateHandler.GetByCPF))))

	mux.Handle("GET /associate/name/{name}", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(associateHandler.GetByName))))

	mux.Handle("PUT /associate/{id}", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(associateHandler.UpdateAssociate))))

	// Rotas Admin
	mux.Handle("POST /admin", http.HandlerFunc(adminHandler.CreateAdmin))

	mux.Handle("POST /login", http.HandlerFunc(adminHandler.Login))

	// Rotas Payment
	mux.Handle("POST /payment", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(paymentHandler.CreatePayment))))

	mux.Handle("GET /payments/month", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(paymentHandler.GetMonthlyTotal))))

	mux.Handle("GET /payments", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(paymentHandler.GetPaymentsByMonth))))

	slog.Info("servidor iniciado", "porta", config.APIPort)
	if err := http.ListenAndServe(config.APIPort, middlewares.CORS(mux)); err != nil {
		slog.Error("erro ao iniciar o servidor", "error", err)
	}

}
