package api

import "net/http"

type Associate interface {
	CreateAssociate(w http.ResponseWriter, r *http.Request)
}
