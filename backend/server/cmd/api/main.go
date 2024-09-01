package main

import (
	"context"
	"log"
	"server/internal/server"
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

	server := server.InitServer(redisClient, dbClient)

	log.Println("Starting server on address", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln("cannot start server: %s", err)
	}
}
