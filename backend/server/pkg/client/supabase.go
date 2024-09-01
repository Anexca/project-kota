package client

import (
	"log"
	"server/pkg/config"

	"github.com/nedpals/supabase-go"
)

func NewSupabaseClient() (*supabase.Client, error) {
	environment, err := config.LoadEnvironment()
	if err != nil {
		return nil, err
	}

	client := supabase.CreateClient(environment.SupabaseUrl, environment.SupabaseKey)
	log.Println("connected to supabase, url:", environment.SupabaseUrl)

	return client, nil
}
