package pkg

import "testing"

func TestCelsiusToKelvinConverter(t *testing.T) {
	tests := []struct {
		celsius  float64
		expected float64
	}{
		{0, 273},
		{100, 373},
		{-273, 0},
	}

	for _, tt := range tests {
		got := CelsiusToKelvinConverter(tt.celsius)
		if got != tt.expected {
			t.Errorf("CelsiusToKelvin(%v) = %v, expected %v", tt.celsius, got, tt.expected)
		}
	}
}
