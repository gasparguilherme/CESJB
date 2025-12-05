package api

import "net/http"

type Associate interface {
	CreateAssociate(w http.ResponseWriter, r *http.Request)
	GetAssociates(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	UpdateAssociate(w http.ResponseWriter, r *http.Request)
}
