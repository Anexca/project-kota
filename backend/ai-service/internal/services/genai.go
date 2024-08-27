package services

import (
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

func (g *GenAIService) GetContentStream(ctx context.Context, prompt, modelName string) (string, error) {
	model := g.client.GenerativeModel(modelName)

	iter := model.GenerateContentStream(ctx, genai.Text(prompt))

	var completeResponse string

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			return "", fmt.Errorf("empty response from model")
		}
		if err != nil {
			return "", err
		}
		for _, c := range resp.Candidates {
			for _, p := range c.Content.Parts {
				completeResponse += fmt.Sprintf("%v", p) // Combine all parts into a single string
			}
		}
	}

	return completeResponse, nil

}
