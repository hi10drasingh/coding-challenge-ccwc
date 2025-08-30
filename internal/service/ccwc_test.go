package service

import (
	"testing"
)

func TestProcess(t *testing.T) {
	content := []byte("Hello, World!\nThis is a test file.\n")
	tests := []struct {
		name    string
		options Options
		output  Results
	}{
		{"count bytes", Options{CountBytes: true}, Results{Bytes: 35}},
		{"count lines", Options{CountLines: true}, Results{Lines: 2}},
		{"count words", Options{CountWords: true}, Results{Words: 7}},
		{"count characters", Options{CountCharacters: true}, Results{Characters: 35}},
		{"count all", Options{}, Results{
			Bytes:      35,
			Lines:      2,
			Words:      7,
			Characters: 35,
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Process(content, tt.options)
			if got != tt.output {
				t.Errorf("Process() = %v; want %v", got, tt.output)
			}
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
		{"emoji", "😊", 4},        // Emoji is 4 bytes in UTF-8
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

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"normal string", "Hello, World!", 13},
		{"empty string", "", 0},
		{"whitespace", "   ", 3},
		{"unicode", "こんにちは", 5},
		{"emoji", "Hello 😊", 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountCharacters(tt.input)
			if got != tt.want {
				t.Errorf("CountCharacters(%q) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
