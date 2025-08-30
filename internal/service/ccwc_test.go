package service

import "testing"

func TestCustomWCTool(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"normal name", "Go", "Custom WC Tool!"},
		{"empty name", "", "Custom WC Tool!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomWCTool()
			if got != tt.want {
				t.Errorf("Greet(%q) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}
