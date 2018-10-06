package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	const src = `
		create_parking_lot 1
		park KA-01-HH-1234 White
		leave 1
		registration_numbers_for_cars_with_colour White
		slot_numbers_for_cars_with_colour White
		slot_number_for_registration_number KA-01-HH-3141
		status
	`

	program, err := Parse(src)
	if err != nil {
		t.Fatalf("parse fail:\n%s", err)
	}

	if l := len(program.Statements); l != 7 {
		t.Fatalf("parse invalid number of statements - want: %d, got: %d", 7, l)
	}
}

func TestParserErrors(t *testing.T) {
	tests := []struct {
		src string
	}{
		{";"},
		{"create_parking_lot 9223372036854775808"}, /* int64 overflow */
		{"create_parking_lot -1"},
		{"park 1 White"},
		{"park KA-01-HH-1111 1"},
		{"leave parking"},
		{"leave 9223372036854775808"}, /* int64 overflow */
		{"registration_numbers_for_cars_with_colour 0"},
		{"slot_numbers_for_cars_with_colour 0"},
		{"slot_number_for_registration_number 0"},
	}

	for _, tt := range tests {
		if _, err := Parse(tt.src); err == nil {
			t.Fatalf("expected error parse but got %q", err)
		}
	}
}
