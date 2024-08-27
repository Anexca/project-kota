package main

import (
	"ai-service/internal/server"
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

	server := server.InitServer(genAiClient, redisClient)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
