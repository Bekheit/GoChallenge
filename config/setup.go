package config

import (
	"context"
	"github.com/jackc/pgx/v4"
)

func ConnectDB() *pgx.Conn {
	conn, _ := pgx.Connect(context.Background(), cockroachURI())
	return conn
}
