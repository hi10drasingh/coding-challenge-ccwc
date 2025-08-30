package service

import (
	"os"
	"testing"
)

func TestProcess(t *testing.T) {
	// Since Process function reads from a file, we will create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up

	content := "Hello, World!"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	tests := []struct {
		name    string
		options map[string]bool
	}{
		{"count bytes", map[string]bool{"-c": true}},
		{"no options", map[string]bool{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Process(tmpFile.Name(), tt.options)
			// Note: In a real test, we would capture stdout and verify the output.
			// For simplicity, we are just ensuring no panic occurs.
		})
	}
}

func TestCountBytes(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"normal string", "Hello, World!", 13},
		{"empty string", "", 0},
		{"whitespace", "   ", 3},
		{"unicode", "こんにちは", 15}, // Each Japanese character is 3 bytes in UTF-8
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountBytes(tt.input)
			if got != tt.want {
				t.Errorf("CountBytes(%q) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
