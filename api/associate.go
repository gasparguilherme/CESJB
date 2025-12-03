package api

import (
	service "cesjb/domain/service/associate"
	handler "cesjb/handlers/associate"
	associated "cesjb/repository/postgresql/associate"

	"github.com/jackc/pgx/v5"
)

func InitAssociated(conn *pgx.Conn) handler.Handler {
	r := associated.NewPostgresRepository(conn)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	return h
}
