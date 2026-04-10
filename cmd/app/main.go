package main

import (
	"cesjb/api"
	"cesjb/config"
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config.Load()

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, config.StringConnectionBase)
	if err != nil {
		slog.Error("Erro ao criar pool de conexões", "error", err.Error())
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		slog.Error("Erro ao fazer ping no banco de dados", "error", err.Error())
		os.Exit(1)
	}

	slog.Info("Conexão estabelecida com sucesso")

	associateHandler := api.InitAssociate(pool)
	adminHandler := api.InitAdmin(pool)
	paymentHandler := api.InitPayment(pool)

	api.StartApp(associateHandler, adminHandler, paymentHandler)
}
