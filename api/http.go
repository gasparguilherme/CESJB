package api

import "net/http"

func StartApp(associateHandler, listAssociatesHandler, getAssociateByIDHandler Associate) {
	mux := http.NewServeMux()

	mux.Handle("POST /associate", http.HandlerFunc(associateHandler.CreateAssociate))
	mux.Handle("GET /associates", http.HandlerFunc(listAssociatesHandler.GetAssociates))
	mux.Handle("GET /associate/{id}", http.HandlerFunc(getAssociateByIDHandler.GetByID))

	http.ListenAndServe(":8088", mux)

}
