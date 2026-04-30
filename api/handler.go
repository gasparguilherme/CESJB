package api

import "net/http"

type Associate interface {
	CreateAssociate(w http.ResponseWriter, r *http.Request)
	GetAssociates(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	GetByCPF(w http.ResponseWriter, r *http.Request)
	GetByName(w http.ResponseWriter, r *http.Request)
	UpdateAssociate(w http.ResponseWriter, r *http.Request)
	GetInactiveAssociates(w http.ResponseWriter, r *http.Request)
}

type Admin interface {
	CreateAdmin(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type Payment interface {
	CreatePayment(w http.ResponseWriter, r *http.Request)
	GetMonthlyTotal(w http.ResponseWriter, r *http.Request)
	GetPaymentsByMonth(w http.ResponseWriter, r *http.Request)
	GetDefaultersByMonth(w http.ResponseWriter, r *http.Request)
	GetPaymentsByAssociate(w http.ResponseWriter, r *http.Request)
	UpdatePayment(w http.ResponseWriter, r *http.Request)
}
