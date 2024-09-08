package client

import (
	"common/ent"
	"context"
	"database/sql"
	"fmt"
	"log"
	"server/pkg/config"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDbClient(ctx context.Context) (*ent.Client, error) {
	environment, err := config.LoadEnvironment()
	if err != nil {
		return nil, err
	}

	databaseUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", environment.DatabaseUser, environment.DatabasePassword, environment.DatabaseHost, environment.DatabasePort, environment.DatabaseName)

	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	log.Println("connected to database server", environment.DatabaseHost)
	return client, nil
}
