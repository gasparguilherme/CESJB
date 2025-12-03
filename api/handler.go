package api

import "net/http"

type Associated interface {
	CreateAssociate(w http.ResponseWriter, r *http.Request)
}
