package client

import (
	"ai-service/pkg/config"
	"common/ent"
	"context"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDbClient(c context.Context, environment *config.Environment) (*ent.Client, error) {
	client, err := ent.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		environment.DatabaseHost, environment.DatabasePort, environment.DatabaseUser, environment.DatabaseName, environment.DatabasePassword,
	))

	if err != nil {
		return nil, err
	}

	err = client.Schema.Create(c)
	return client, err
}
