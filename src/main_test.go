package main

import (
	"os"
	"testing"
)

func TestMainFunction(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestMainFunction"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestGetAWSToken(t *testing.T) {
	os.Setenv("AWS_TOKEN", "dummy_token")
	defer os.Unsetenv("AWS_TOKEN")

	expectedToken := "dummy_token"
	token := GetAWSToken()

	if token != expectedToken {
		t.Errorf("Expected: %s, Got: %s", expectedToken, token)
	}
}
