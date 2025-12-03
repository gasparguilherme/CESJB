package api

import "net/http"

func StartApp(associateHandler Associate) {
	mux := http.NewServeMux()

	mux.Handle("POST /associate", http.HandlerFunc(associateHandler.CreateAssociate))

	http.ListenAndServe(":8088", mux)

}
