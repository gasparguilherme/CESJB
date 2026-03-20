package main

import (
	"cesjb/api"
	"cesjb/config"
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	config.Load()

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, config.StringConnectionBase)
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
	paymentAssociateHandler := api.InitAssociate(conn)

	//Admin
	adminHandler := api.InitAdmin(conn)
	loginHandler := api.InitAdmin(conn)

	//start api
	api.StartApp(associateHandler, listHandler, getByIDHandler, updateAssociateHandler,
		paymentAssociateHandler, adminHandler, loginHandler)

}
