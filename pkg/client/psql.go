package client

import (
	"fmt"
	"time"

	"github.com/andy-ahmedov/makves/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDB(cfg config.Postgres) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
