package database

import "testing"

func TestNewCar(t *testing.T) {
	tests := []struct {
		registrationNumber string
		color              string
		expectedError      bool
	}{
		{"AA-00-AA-0000", "White", false},
		{"AA-00-A-0000", "White", false},
		{"AA-00-AA-000", "White", false},
		{"AA-00-A-000", "White", false},

		{"AA-00-AA-0000", "White", false},
		{"AA-00-AA-0000", "Yellow", false},
		{"AA-00-AA-0000", "Red", false},
		{"AA-00-AA-0000", "Blue", false},
		{"AA-00-AA-0000", "Green", false},
		{"AA-00-AA-0000", "Brown", false},
		{"AA-00-AA-0000", "Black", false},

		{"aa-00-aa-0000", "White", true},
		{"A-00-AA-0000", "White", true},
		{"AA-0-AA-0000", "White", true},
		{"AA-00-AA-00000", "White", true},
		{"AA-00-AA-0000", "Invalid", true},
	}

	for _, tt := range tests {
		_, err := NewCar(tt.registrationNumber, tt.color)
		if tt.expectedError && err == nil {
			t.Errorf("car(%s, %s) expected error but got: <nil>", tt.registrationNumber, tt.color)
		}
		if !tt.expectedError && err != nil {
			t.Errorf("car(%s, %s) expected no error but got: %s", tt.registrationNumber, tt.color, err)
		}
	}
}
