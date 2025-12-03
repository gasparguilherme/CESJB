package app

import "net/http"

type Associated interface {
	CreateAssociated(w http.ResponseWriter, r *http.Request)
}
