package api

import (
	service "cesjb/domain/service/admin"
	handler "cesjb/handlers/admin"
	admin "cesjb/repository/postgresql/admin"

	"github.com/jackc/pgx/v5"
)

func InitAdmin(conn *pgx.Conn) handler.Handler {
	r := admin.NewPostgresRepository(conn)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	return h
}
