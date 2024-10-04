package mocks

import (
	"common/constants"
	"context"
	"time"

	"github.com/stretchr/testify/mock"
)

// Mock for GenAIServiceInterface
type MockGenAIService struct {
	mock.Mock
}

func (m *MockGenAIService) GetContentStream(ctx context.Context, prompt string, model constants.GenAiModel) (string, error) {
	args := m.Called(ctx, prompt, model)
	return args.String(0), args.Error(1)
}

func (m *MockGenAIService) GetStructuredContentStream(ctx context.Context, prompt string, modelName constants.GenAiModel) (string, error) {
	args := m.Called(ctx, prompt, modelName)
	return args.String(0), args.Error(1)
}

// Mock for RedisServiceInterface
type MockRedisService struct {
	mock.Mock
}

// Mock implementation for Store method
func (m *MockRedisService) Store(ctx context.Context, key string, value string, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

// Mock implementation for Get method
func (m *MockRedisService) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

// Mock implementation for Delete method
func (m *MockRedisService) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

// Mock implementation for Health method
func (m *MockRedisService) Health() map[string]string {
	args := m.Called()
	return args.Get(0).(map[string]string)
}

// Mock implementation for CheckRedisHealth method
func (m *MockRedisService) CheckRedisHealth(ctx context.Context, stats map[string]string) map[string]string {
	args := m.Called(ctx, stats)
	return args.Get(0).(map[string]string)
}

// Mock for ProfanityServiceInterface
type MockProfanityService struct {
	mock.Mock
}

func (m *MockProfanityService) IsProfane(s string) bool {
	args := m.Called(s)
	return args.Bool(0)
}
