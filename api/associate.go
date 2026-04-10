package api

import (
	service "cesjb/domain/service/associate"
	handler "cesjb/handlers/associate"
	associate "cesjb/repository/postgresql/associate"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitAssociate(pool *pgxpool.Pool) handler.Handler {
	r := associate.NewPostgresRepository(pool)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	return h
}
