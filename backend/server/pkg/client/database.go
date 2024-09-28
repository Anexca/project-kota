package client

import (
	"context"
	"fmt"
	"log"

	"common/ent"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"server/pkg/config"
)

func NewDbClient(ctx context.Context) (*ent.Client, error) {
	environment, err := config.LoadEnvironment()
	if err != nil {
		return nil, err
	}

	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", environment.DatabaseUser, environment.DatabasePassword, environment.DatabaseHost, environment.DatabasePort, environment.DatabaseName)

	poolConfig, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	poolConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatal(err)
	}

	db := stdlib.OpenDBFromPool(pool)
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	log.Println("connected to database server", environment.DatabaseHost)
	return client, nil
}
