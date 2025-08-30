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

	content := "Hello, World!\nThis is a test file.\n"
	if _, err := tmpFile.WriteString(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	tests := []struct {
		name    string
		options map[string]bool
	}{
		{"count bytes", map[string]bool{"-c": true}},
		{"count lines", map[string]bool{"-l": true}},
		{"count words", map[string]bool{"-w": true}},
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

func TestCountLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"multiple lines", "Hello\nWorld\nThis is a test.", 3},
		{"single line", "Hello, World!", 1},
		{"empty string", "", 0},
		{"empty line", "\n", 1},
		{"trailing newline", "Hello, World!\n", 1},
		{"only newlines", "\n\n\n", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountLines(tt.input)
			if got != tt.want {
				t.Errorf("CountLines(%q) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"normal sentence", "Hello, World! This is a test.", 6},
		{"multiple spaces", "Hello,   World!  This   is a test.", 6},
		{"newlines and tabs", "Hello,\nWorld!\tThis is a test.", 6},
		{"empty string", "", 0},
		{"only spaces", "     ", 0},
		{"only newlines", "\n\n\n", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.input)
			if got != tt.want {
				t.Errorf("CountWords(%q) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
