package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	commonConstants "common/constants"

	"server/pkg/config"
)

type PromptService struct {
}

func NewPromptService() *PromptService {
	return &PromptService{}
}

func (p *PromptService) GetPromptResult(ctx context.Context, prompt string, model commonConstants.GenAiModel) (string, error) {
	env, err := config.LoadEnvironment()
	if err != nil {
		return "", fmt.Errorf("failed to load environment: %v", err)
	}

	promptUrl := fmt.Sprintf("%s/prompt", env.AIServiceUrl)

	req, err := prepareRequest(ctx, promptUrl, env.AIServiceAccessKey, prompt, model)
	if err != nil {
		return "", err
	}

	responseBody, err := sendRequest(req)
	if err != nil {
		return "", err
	}

	return responseBody, nil
}

func (p *PromptService) PingServer(ctx context.Context) error {
	env, err := config.LoadEnvironment()
	if err != nil {
		return fmt.Errorf("failed to load environment: %v", err)
	}

	supRequestUrl := fmt.Sprintf("%s/sup", env.AIServiceUrl)

	req, err := http.NewRequestWithContext(ctx, "GET", supRequestUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	_, err = sendRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func prepareRequest(ctx context.Context, url, accessKey, prompt string, model commonConstants.GenAiModel) (*http.Request, error) {
	postData := map[string]string{
		"prompt": prompt,
		"model":  string(model),
	}

	jsonData, err := json.Marshal(postData)
	if err != nil {
		return nil, fmt.Errorf("error encoding JSON: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", accessKey))

	return req, nil
}

func sendRequest(req *http.Request) (string, error) {
	client := &http.Client{
		Timeout: 2 * time.Minute,
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
