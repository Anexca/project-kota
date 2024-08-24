package main

import (
	"ai-service/internal/server"
	"ai-service/pkg/client"
	"context"
	"fmt"
	"log"
)

func main() {
	// err := client.MakeChatRequests(context.Background(), "project-kota-433508", "asia-east1", "gemini-1.5-flash")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	ctx := context.Background()

	genAiClient, err := client.NewGeminiClient(ctx, "project-kota-433508", "asia-east1")
	if err != nil {
		log.Fatalln(err)
	}
	defer genAiClient.Close()

	server := server.InitServer(genAiClient)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}

}
