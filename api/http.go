package api

import (
	"fmt"
	"net/http"
)

func StartApp(associateHandler, listAssociatesHandler, getAssociateByIDHandler, updateAssociateHandler,
	getAssociateByCPF Associate) {
	mux := http.NewServeMux()

	mux.Handle("POST /associate", http.HandlerFunc(associateHandler.CreateAssociate))
	mux.Handle("GET /associates", http.HandlerFunc(listAssociatesHandler.GetAssociates))
	mux.Handle("GET /associate/id/{id}", http.HandlerFunc(getAssociateByIDHandler.GetByID))
	mux.Handle("PUT /associate/{id}", http.HandlerFunc(updateAssociateHandler.UpdateAssociate))
	mux.Handle("GET /associate/cpf/{cpf}", http.HandlerFunc(getAssociateByCPF.GetAssociateByCPF))
	fmt.Println("Rota CPF registrada!")

	http.ListenAndServe(":8088", mux)

}
