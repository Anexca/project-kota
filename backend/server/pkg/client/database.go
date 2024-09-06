package client

import (
	"common/ent"
	"context"
	"fmt"
	"log"
	"server/pkg/config"

	_ "github.com/lib/pq"
)

func NewDbClient(ctx context.Context) (*ent.Client, error) {
	environment, err := config.LoadEnvironment()
	if err != nil {
		return nil, err
	}

	client, err := ent.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s binary_parameters=yes",
		environment.DatabaseHost, environment.DatabasePort, environment.DatabaseUser, environment.DatabaseName, environment.DatabasePassword,
	))

	if err != nil {
		return nil, err
	}

	err = client.Schema.Create(ctx)
	log.Println("connected to database server", environment.DatabaseHost)
	return client, err
}
