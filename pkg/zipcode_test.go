package pkg

import "testing"

func TestIsZipcodeValid(t *testing.T) {
	tests := []struct {
		cep      string
		expected bool
	}{
		{"78590-000", true},
		{"12345-678", true},
		{"12345", false},
		{"78590000", true},
		{"test", false},
		{"78S90000", false},
	}

	for _, tt := range tests {
		t.Run(tt.cep, func(t *testing.T) {
			actual := IsZipcodeValid(tt.cep)
			if actual != tt.expected {
				t.Errorf("IsValidCEP(%v) = %v; expected %v", tt.cep, actual, tt.expected)
			}
		})
	}
}
