package api

import "net/http"

func StartApp(associatedHandler Associated) {
	mux := http.NewServeMux()

	mux.Handle("POST /associated", http.HandlerFunc(associatedHandler.CreateAssociated))

	http.ListenAndServe(":8088", mux)

}
