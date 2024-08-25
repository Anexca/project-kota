package client

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/vertexai/genai"
)

func NewGeminiClient(ctx context.Context, projectID, region string) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, projectID, region)
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %w", err)
	}

	log.Println("GENAI client connected")

	return client, nil
}
