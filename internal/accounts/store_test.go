package accounts

import (
	"testing"
)

func TestNormalizeEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "already normalized",
			input:    "mabel@example.com",
			expected: "mabel@example.com",
		},
		{
			name:     "uppercase",
			input:    "MABEL@EXAMPLE.COM",
			expected: "mabel@example.com",
		},
		{
			name:     "mixed case with surrounding whitespace",
			input:    "  Ligma@Example.COM \t\n ",
			expected: "ligma@example.com",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only whitespace",
			input:    "   ",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeEmail(tt.input)
			if got != tt.expected {
				t.Errorf("NormalizeEmail(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}
