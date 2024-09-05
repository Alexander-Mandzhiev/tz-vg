package postgres

import (
	"context"
	"fmt"
	"tz-vg/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresDB(cfg config.Database) (*pgxpool.Pool, error) {
	databaseString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	pool, err := pgxpool.New(context.Background(), databaseString)
	return pool, err
}
