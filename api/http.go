package api

import "net/http"

func StartApp(associateHandler, getAssociatesHandler, get_By_ID_Handler Associate) {
	mux := http.NewServeMux()

	mux.Handle("POST /associate", http.HandlerFunc(associateHandler.CreateAssociate))
	mux.Handle("GET /associates", http.HandlerFunc(getAssociatesHandler.GetAssociates))
	mux.Handle("GET /associate/{id}", http.HandlerFunc(get_By_ID_Handler.GetByID))

	http.ListenAndServe(":8088", mux)

}
