package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"gitlab.com/tantai-smap/authenticate-api/config"
)

const (
	connectTimeout = 10 * time.Second
)

func Connect(ctx context.Context, cfg config.PostgresConfig) (*sql.DB, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), connectTimeout)
	defer cancelFunc()

	print(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName))
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping to DB: %w", err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(time.Minute)

	log.Println("Connected to Postgres!")

	return db, nil
}

func Disconnect(ctx context.Context, db *sql.DB) error {
	if db == nil {
		return nil
	}

	if err := db.Close(); err != nil {
		return fmt.Errorf("failed to close DB: %w", err)
	}

	log.Println("Disconnected from Postgres!")

	return nil
}
