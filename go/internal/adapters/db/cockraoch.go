package db

import (
	"example/go/config"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"go.elastic.co/apm/module/apmsql/v2"
	"go.uber.org/zap"
)

func NewDatabaseConnection(log *zap.SugaredLogger, config config.DatabaseConfigurations) *bun.DB {
	apmsql.Register("postgres", &stdlib.Driver{})
	sqlDb, err := apmsql.Open("postgres", config.Dsn)

	if err != nil {
		log.Fatalf("DB connection error -> %v", err)
	}

	sqlDb.SetMaxOpenConns(config.Pool)

	err = sqlDb.Ping()
	if err != nil {
		log.Fatalf("DB connection error -> %v", err)
	}

	db := bun.NewDB(sqlDb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())

	log.Infof("Database connected successfully. Connections opened: %d", db.Stats().OpenConnections)

	return db
}
