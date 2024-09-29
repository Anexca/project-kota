package services_test

import (
	"ai-service/internal/services"
	"context"
	"testing"

	"common/constants"

	"cloud.google.com/go/vertexai/genai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/api/iterator"
)

// Mock the GenAIClient interface
type MockGenAIClient struct {
	mock.Mock
}

func (m *MockGenAIClient) GenerativeModel(modelName string) services.GenAIGenerativeModel {
	args := m.Called(modelName)
	return args.Get(0).(services.GenAIGenerativeModel)
}

// Mock the GenAIGenerativeModel interface
type MockGenAIGenerativeModel struct {
	mock.Mock
}

func (m *MockGenAIGenerativeModel) GenerateContentStream(ctx context.Context, text genai.Text) services.GenAIContentIterator {
	args := m.Called(ctx, text)
	return args.Get(0).(services.GenAIContentIterator)
}

// Mock the GenAIContentIterator interface
type MockGenAIContentIterator struct {
	mock.Mock
}

func (m *MockGenAIContentIterator) Next() (*genai.GenerateContentResponse, error) {
	args := m.Called()
	return args.Get(0).(*genai.GenerateContentResponse), args.Error(1)
}

func TestGetContentStream_Success(t *testing.T) {
	// Arrange
	mockClient := new(MockGenAIClient)
	mockModel := new(MockGenAIGenerativeModel)
	mockIterator := new(MockGenAIContentIterator)

	service := services.NewGenAIService(mockClient)

	// Mock behavior
	mockClient.On("GenerativeModel", "test-model").Return(mockModel)
	mockModel.On("GenerateContentStream", mock.Anything, mock.Anything).Return(mockIterator)

	resp := &genai.GenerateContentResponse{
		Candidates: []*genai.Candidate{
			{
				Content: &genai.Content{
					Parts: []genai.Part{"Hello", " World"}, // Ensure genai.Part is mocked correctly
				},
			},
		},
	}
	mockIterator.On("Next").Return(resp, nil).Once()
	mockIterator.On("Next").Return(nil, iterator.Done)

	// Act
	result, err := service.GetContentStream(context.Background(), "test prompt", constants.GenAiModel("test-model"))

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", result)
}
