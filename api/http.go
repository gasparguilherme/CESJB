package api

import (
	"cesjb/middlewares"
	"net/http"
)

func StartApp(associateHandler, listAssociatesHandler, getAssociateByIDHandler, updateAssociateHandler,
	paymentAssociateHandler Associate, adminHandler, loginHandler Admin) {
	mux := http.NewServeMux()

	// Rotas Associado
	mux.Handle("POST /associate", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(associateHandler.CreateAssociate))))

	mux.Handle("GET /associates", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(listAssociatesHandler.GetAssociates))))

	mux.Handle("GET /associate/id/{id}", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(getAssociateByIDHandler.GetByID))))

	mux.Handle("PUT /associate/{id}", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(updateAssociateHandler.UpdateAssociate))))

	//Rotas Admin
	mux.Handle("POST /admin", http.HandlerFunc(adminHandler.CreateAdmin))

	mux.Handle("POST /login", http.HandlerFunc(loginHandler.Login))

	//Rotas Payment
	mux.Handle("POST /payment", middlewares.Logger(middlewares.Authenticate(
		http.HandlerFunc(paymentAssociateHandler.CreatePayment))))

	http.ListenAndServe(":8088", mux)

}
