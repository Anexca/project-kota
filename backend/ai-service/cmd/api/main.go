package main

import (
	"ai-service/internal/server"
	"ai-service/internal/workers"
	"ai-service/pkg/client"
	"context"
	"fmt"
	"log"
)

func main() {
	ctx := context.Background()

	genAiClient, err := client.NewGeminiClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer genAiClient.Close()

	redisClient, err := client.NewRedisClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer redisClient.Close()

	c := workers.InitWorkers(genAiClient, redisClient)
	c.Start()

	defer c.Stop()

	server := server.InitServer(genAiClient, redisClient)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(fmt.Sprintf("cannot start server: %s", err))
	}

}
