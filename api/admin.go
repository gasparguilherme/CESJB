package api

import (
	service "cesjb/domain/service/admin"
	handler "cesjb/handlers/admin"
	admin "cesjb/repository/postgresql/admin"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitAdmin(pool *pgxpool.Pool) handler.Handler {
	r := admin.NewPostgresRepository(pool)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	return h
}
