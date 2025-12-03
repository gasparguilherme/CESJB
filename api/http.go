package api

import "net/http"

func StartApp(associateHandler Associate, getAssociateHandler Associate) {
	mux := http.NewServeMux()

	mux.Handle("POST /associate", http.HandlerFunc(associateHandler.CreateAssociate))
	mux.Handle("GET /associate", http.HandlerFunc(associateHandler.GetAssociates))

	http.ListenAndServe(":8088", mux)

}
