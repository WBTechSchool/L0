package db

import (
	"context"
	"os"

	"github.com/WBTechSchool/L0/internal/gen"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	Queries *gen.Queries
	Pool    *pgxpool.Pool
}

func New(ctx context.Context) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("BACKEND_DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	queries := gen.New(pool)

	return &Storage{Queries: queries, Pool: pool}, nil
}
