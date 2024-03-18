package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

// Connect to Postgres with func params
func NewConnection(ctx context.Context, url string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к Postgres: %+v\n", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("Ошибка проверки подключения к Postgres: %+v", err)
	}
	return conn, nil
}

// NewPoolConnection new connection to pool
func NewPoolConnection(ctx context.Context, url string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("Не удалось подключиться к Postgres Pool: %+v\n", err)
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("Ошибка проверки подключения к Postgres: %+v", err)
	}
	return pool, nil
}

func Close(conn *pgx.Conn, ctx context.Context) {
	err := conn.Close(ctx)
	if err != nil {
		log.Fatalf("Ошибка отключения соединения с Postgres: %+v", err)
	}
}
