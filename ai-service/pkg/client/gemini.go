package client

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/vertexai/genai"
	"google.golang.org/api/iterator"
)

func MakeChatRequests(ctx context.Context, projectID, region, modelName string) error {
	client, err := genai.NewClient(ctx, projectID, region)
	if err != nil {
		return fmt.Errorf("unable to create client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)

	iter := model.GenerateContentStream(
		ctx,
		genai.Text(`Generate a JSON array containing 5 multiple-choice questions (MCQs) for the Chemistry subject in JEE mains exam. 
			Each MCQ should have a "question", "options", "answer", and "explanation" field. 
			The JSON output should be formatted as a single-line string without any extra whitespace or formatting`),
	)

	var completeResponse string

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
			return fmt.Errorf("empty response from model")
		}
		if err != nil {
			return err
		}
		for _, c := range resp.Candidates {
			for _, p := range c.Content.Parts {
				completeResponse += fmt.Sprintf("%v", p) // Combine all parts into a single string
			}
		}
	}

	log.Println(completeResponse)

	return nil
}

func NewGeminiClient(ctx context.Context, projectID, region string) (*genai.Client, error) {
	client, err := genai.NewClient(ctx, projectID, region)
	if err != nil {
		return nil, fmt.Errorf("unable to create client: %w", err)
	}

	return client, nil
}
