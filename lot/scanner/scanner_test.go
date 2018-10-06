package scanner

import (
	"testing"

	"parking_lot/lot/token"
)

func TestScanSingleToken(t *testing.T) {
	tests := []struct {
		src string
		tok token.Token
	}{
		{";", token.ILLEGAL},
		{" ", token.EOF},
		{"0", token.INT},
		{"a-", token.STRING},
		{"create_parking_lot", token.CREATE_PARKING_LOT},
		{"park", token.PARK},
		{"leave", token.LEAVE},
		{"status", token.STATUS},
		{"registration_numbers_for_cars_with_colour", token.REGISTRATION_NUMBERS_FOR_CARS_WITH_COLOUR},
		{"slot_numbers_for_cars_with_colour", token.SLOT_NUMBERS_FOR_CARS_WITH_COLOUR},
		{"slot_number_for_registration_number", token.SLOT_NUMBER_FOR_REGISTRATION_NUMBER},
	}

	for _, tt := range tests {
		pos, tok, lit := New(tt.src).Scan()
		if tok != tt.tok {
			t.Fatalf("unexpected token in %q source - want: %s, got: %s", tt.src, tt.tok, tok)
		}

		if expectedLit := tt.src[pos : pos+len(lit)]; lit != expectedLit {
			t.Fatalf("unexpected literal in %q source - want: %s, got: %s", tt.src, expectedLit, lit)
		}
	}
}

func TestScanMultipleToken(t *testing.T) {
	tests := []struct {
		src    string
		tokens []token.Token
	}{
		{
			"create_parking_lot 1",
			[]token.Token{token.CREATE_PARKING_LOT, token.INT},
		},
		{
			"park KA-01-AA-0000 White",
			[]token.Token{token.PARK, token.STRING, token.STRING},
		},
		{
			"status leave 1",
			[]token.Token{token.STATUS, token.LEAVE, token.INT},
		},
		{
			"registration_numbers_for_cars_with_colour White\n\tslot_numbers_for_cars_with_colour White",
			[]token.Token{
				token.REGISTRATION_NUMBERS_FOR_CARS_WITH_COLOUR,
				token.STRING,
				token.SLOT_NUMBERS_FOR_CARS_WITH_COLOUR,
				token.STRING,
			},
		},
	}

	for _, tt := range tests {
		s := New(tt.src)
		for _, expectedToken := range tt.tokens {
			pos, tok, lit := s.Scan()
			if tok != expectedToken {
				t.Fatalf("unexpected token in %q source - want: %s, got: %s", tt.src, expectedToken, tok)
			}

			if expectedLit := tt.src[pos : pos+len(lit)]; lit != expectedLit {
				t.Fatalf("unexpected literal in %q source - want: %s, got: %s", tt.src, expectedLit, lit)
			}
		}
	}
}
