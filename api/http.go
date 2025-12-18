package api

import (
	"net/http"
)

func StartApp(associateHandler, listAssociatesHandler, getAssociateByIDHandler, updateAssociateHandler,
	getAssociateByCPF Associate, adminHandler, loginHandler Admin) {
	mux := http.NewServeMux()
	// Rotas Associado
	mux.Handle("POST /associate", http.HandlerFunc(associateHandler.CreateAssociate))
	mux.Handle("GET /associates", http.HandlerFunc(listAssociatesHandler.GetAssociates))
	mux.Handle("GET /associate/id/{id}", http.HandlerFunc(getAssociateByIDHandler.GetByID))
	mux.Handle("PUT /associate/{id}", http.HandlerFunc(updateAssociateHandler.UpdateAssociate))
	mux.Handle("GET /associate/cpf/{cpf}", http.HandlerFunc(getAssociateByCPF.GetAssociateByCPF))

	//Rotas Admin
	mux.Handle("POST /admin", http.HandlerFunc(adminHandler.CreateAdmin))
	mux.Handle("POST /login", http.HandlerFunc(loginHandler.FindAdminByEmail))

	http.ListenAndServe(":8088", mux)

}
