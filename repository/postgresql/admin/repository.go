package admin

import "github.com/jackc/pgx/v5/pgxpool"

type Repository struct {
	connectionInstance *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) Repository {
	return Repository{
		connectionInstance: pool,
	}
}
