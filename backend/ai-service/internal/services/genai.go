package services

import (
	"context"
	"fmt"

	"cloud.google.com/go/vertexai/genai"
	"google.golang.org/api/iterator"

	commonConstants "common/constants"
)

// Define the interfaces to abstract the external dependencies

// GenAIClient is the interface for the GenAI Client
type GenAIClient interface {
	GenerativeModel(modelName string) GenAIGenerativeModel
}

// GenAIGenerativeModel is the interface for the generative model
type GenAIGenerativeModel interface {
	GenerateContentStream(ctx context.Context, text genai.Text) GenAIContentIterator
}

// GenAIContentIterator is the interface for the content stream iterator
type GenAIContentIterator interface {
	Next() (*genai.GenerateContentResponse, error)
}

// Wrapper for the genai.Client to implement the GenAIClient interface
type GenAIClientWrapper struct {
	client *genai.Client
}

// NewGenAIClientWrapper returns a new wrapper for the genai.Client
func NewGenAIClientWrapper(client *genai.Client) *GenAIClientWrapper {
	return &GenAIClientWrapper{
		client: client,
	}
}

// Implement the GenerativeModel method for the wrapper
func (g *GenAIClientWrapper) GenerativeModel(modelName string) GenAIGenerativeModel {
	return &GenAIGenerativeModelWrapper{
		model: g.client.GenerativeModel(modelName),
	}
}

// Wrapper for genai.GenerativeModel to implement GenAIGenerativeModel interface
type GenAIGenerativeModelWrapper struct {
	model *genai.GenerativeModel
}

// Implement the GenerateContentStream method
func (g *GenAIGenerativeModelWrapper) GenerateContentStream(ctx context.Context, text genai.Text) GenAIContentIterator {
	return &GenAIContentIteratorWrapper{
		iter: g.model.GenerateContentStream(ctx, text),
	}
}

// Wrapper for genai.ContentStreamIterator to implement GenAIContentIterator interface
type GenAIContentIteratorWrapper struct {
	iter *genai.GenerateContentResponseIterator
}

// Implement the Next method for the content iterator
func (g *GenAIContentIteratorWrapper) Next() (*genai.GenerateContentResponse, error) {
	return g.iter.Next()
}

// GenAIService struct with the client as an interface
type GenAIService struct {
	client GenAIClient
}

// NewGenAIService constructor using the interface
func NewGenAIService(client GenAIClient) *GenAIService {
	return &GenAIService{
		client: client,
	}
}

// GetContentStream method uses the GenAIClient interface for generating the content stream
func (g *GenAIService) GetContentStream(ctx context.Context, prompt string, modelName commonConstants.GenAiModel) (string, error) {
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
