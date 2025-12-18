package main

import (
	"cesjb/api"
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	dbURL := "postgres://postgres:senha@localhost:5433/cesjb"

	conn, err := pgx.Connect(ctx, dbURL)

	if err != nil {
		slog.Error("Erro ao conectar no banco", "error", err.Error())
		os.Exit(1)
	}
	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		slog.Error("Error ao fazer ping no banco de dados", "error", err.Error())
		os.Exit(1)
	}

	slog.Info("Conexão estabelcida com sucesso")

	//Associado
	associateHandler := api.InitAssociate(conn)
	listHandler := api.InitAssociate(conn)
	getByIDHandler := api.InitAssociate(conn)
	updateAssociateHandler := api.InitAssociate(conn)
	getAssociateByCPF := api.InitAssociate(conn)

	//Admin
	adminHandler := api.InitAdmin(conn)
	loginHandler := api.InitAdmin(conn)
	api.StartApp(associateHandler, listHandler, getByIDHandler, updateAssociateHandler, getAssociateByCPF,
		adminHandler, loginHandler)

}
