package main

import (
	"context"
	"log"

	"ai-service/internal/server"
	"ai-service/internal/workers"
	"ai-service/pkg/client"
)

func main() {
	ctx := context.Background()

	genAiClient, err := client.NewGeminiClient(ctx)
	if err != nil {
		log.Println(err)
	}
	defer genAiClient.Close()

	redisClient, err := client.NewRedisClient(ctx)
	if err != nil {
		log.Println(err)
	}
	defer redisClient.Close()

	dbclient, err := client.NewDbClient(ctx)
	if err != nil {
		log.Println(err)
	}
	defer dbclient.Close()

	c := workers.InitWorkers(genAiClient, redisClient, dbclient)
	c.Start()

	defer c.Stop()

	server := server.InitServer(genAiClient, redisClient)

	err = server.ListenAndServe()
	if err != nil {
		log.Printf("cannot start server: %s", err)
	}

}
