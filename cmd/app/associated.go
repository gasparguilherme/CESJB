package main

import (
	service "cesjb/domain/service/associated"
	handler "cesjb/handlers/associated"
	"cesjb/repository/postgresql/associated"

	"github.com/jackc/pgx/v5"
)

func InitAssociated(conn *pgx.Conn) handler.Handler {
	r := associated.NewPostgresRepository(conn)
	s := service.NewService(r)
	h := handler.NewHandler(s)
	return h
}
