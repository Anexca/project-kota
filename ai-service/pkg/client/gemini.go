package client

import (
	"ai-service/pkg/config"
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/vertexai/genai"
)

func NewGeminiClient(ctx context.Context) (*genai.Client, error) {
	env, err := config.LoadEnvironment()
	if err != nil {
		return nil, err
	}

	client, err := genai.NewClient(ctx, env.GoogleCloudProjectId, env.GoogleCloudProjectRegion)
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %w", err)
	}

	log.Println("GENAI client connected")

	return client, nil
}
