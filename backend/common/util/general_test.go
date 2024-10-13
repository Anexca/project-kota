package util_test

import (
	"common/util"
	"testing"

	"github.com/google/uuid"
)

func TestGenerateUUID(t *testing.T) {
	generatedUUID := util.GenerateUUID()

	_, err := uuid.Parse(generatedUUID)
	if err != nil {
		t.Errorf("Expected valid UUID, but got an error: %v", err)
	}

	if generatedUUID == "" {
		t.Errorf("Expected a non-empty UUID, but got an empty string")
	}
}
