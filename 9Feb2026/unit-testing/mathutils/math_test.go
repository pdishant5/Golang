package mathutils

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"both positive", 2, 3, 5},
		{"with zero", 0, 5, 5},
		{"both negative", -2, -3, -5},
		{"mixed signs", -2, 3, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Expected: %d, Got: %d", tt.expected, result)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"even number", 4, true},
		{"odd number", 7, false},
		{"zero", 0, true},
		{"negative even", -6, true},
		{"negative odd", -3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsEven(tt.input)
			if result != tt.expected {
				t.Errorf("Expected: %v, Got: %v", tt.expected, result)
			}
		})
	}
}
