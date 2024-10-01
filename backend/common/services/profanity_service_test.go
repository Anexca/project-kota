package services_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"common/services"
)

func TestProfanityService_IsProfane(t *testing.T) {
	profanityService := services.NewProfanityService()

	tests := []struct {
		input          string
		expectedResult bool
	}{
		{"This is a clean sentence.", false},
		{"This is a bad word shit!", true},
		{"Another sentence with profane language fucker!", true},
		{"No swearing here.", false},
		{"This is just an example.", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := profanityService.IsProfane(test.input)
			require.Equal(t, test.expectedResult, result)
		})
	}
}

func TestProfanityService_WithCustomDictionaries(t *testing.T) {
	profanityService := services.NewProfanityService()

	// Example tests with your custom dictionaries
	tests := []struct {
		input          string
		expectedResult bool
	}{
		{"Explicit bad fuck!", true},       // Example of a phrase that should be caught
		{"Just a normal sentence.", false}, // Should not be caught
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := profanityService.IsProfane(test.input)
			require.Equal(t, test.expectedResult, result)
		})
	}
}
