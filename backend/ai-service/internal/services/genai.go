package services

import (
	"ai-service/pkg/constants"
	"context"
	"fmt"

	"cloud.google.com/go/vertexai/genai"
	"google.golang.org/api/iterator"
)

type GenAIService struct {
	client *genai.Client
}

func NewGenAIService(client *genai.Client) *GenAIService {
	return &GenAIService{
		client: client,
	}
}

func (g *GenAIService) GetContentStream(ctx context.Context, prompt string, modelName constants.GenAiModel) (string, error) {
	if prompt == "" {
		return "", fmt.Errorf("prompt cannot be empty")
	}
	if modelName == "" {
		return "", fmt.Errorf("model name cannot be empty")
	}

	model := g.client.GenerativeModel(string(modelName))
	if model == nil {
		return "", fmt.Errorf("failed to initialize generative model: %s", modelName)
	}

	iter := model.GenerateContentStream(ctx, genai.Text(prompt))
	if iter == nil {
		return "", fmt.Errorf("failed to create content stream iterator")
	}

	var completeResponse string

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return "", fmt.Errorf("error retrieving next content: %w", err)
		}

		if resp == nil || len(resp.Candidates) == 0 {
			return "", fmt.Errorf("empty or invalid response from model")
		}

		for _, c := range resp.Candidates {
			if c.Content == nil || len(c.Content.Parts) == 0 {
				continue // Skip candidates with no content parts
			}
			for _, p := range c.Content.Parts {
				completeResponse += fmt.Sprintf("%v", p)
			}
		}
	}

	if completeResponse == "" {
		return "", fmt.Errorf("no content generated from model")
	}

	return completeResponse, nil
}
