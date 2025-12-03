package api

import "net/http"

func StartApp(associatedHandler Associated) {
	mux := http.NewServeMux()

	mux.Handle("POST /associate", http.HandlerFunc(associatedHandler.CreateAssociate))

	http.ListenAndServe(":8088", mux)

}
