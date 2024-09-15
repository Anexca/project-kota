package main

import (
	"context"
	"log"
	"server/internal/server"
	"server/internal/workers"
	"server/pkg/client"
)

func main() {
	ctx := context.Background()

	redisClient, err := client.NewRedisClient(ctx)
	if err != nil {
		log.Fatalln("cannot connect to redis", err)
	}
	defer redisClient.Close()

	dbClient, err := client.NewDbClient(ctx)
	if err != nil {
		log.Fatalln("cannot connect to database", err)
	}
	defer dbClient.Close()

	supabaseClient, err := client.NewSupabaseClient()
	if err != nil {
		log.Fatalln("cannot connect to supabase", err)
	}

	paymentClient, err := client.NewRazorpayClient()
	if err != nil {
		log.Fatalln("cannot connect to payment client", err)
	}

	workers := workers.InitWorkers(redisClient, dbClient)
	defer workers.Stop()

	server := server.InitServer(redisClient, dbClient, supabaseClient, paymentClient)

	log.Println("Starting server on address", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
