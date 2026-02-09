package stringutils

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple word", "go", "og"},
		{"palindrome", "madam", "madam"},
		{"with spaces", "go lang", "gnal og"},
		{"unicode1", "नमस्ते", "ेत्समन"},
		{"unicode2", "મલયાલમ", "મલાયલમ"},
		{"empty string", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Expected: %q; Got: %q", tt.expected, result)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"simple palindrome", "madam", true},
		{"case insensitive", "Level", true},
		{"not palindrome", "golang", false},
		{"single character", "a", true},
		{"empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("Expected: %v; Got: %v", tt.expected, result)
			}
		})
	}
}
