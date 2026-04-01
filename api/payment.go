package api

import (
	service "cesjb/domain/service/payment"
	handler "cesjb/handlers/payment"
	associated "cesjb/repository/postgresql/payment"

	"github.com/jackc/pgx/v5"
)

func InitPayment(conn *pgx.Conn) handler.Handler {
	r := associated.NewPostgresRepository(conn)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	return h
}
