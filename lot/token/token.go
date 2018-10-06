// Package token defines constants representing the lexical tokens of the LoT language.
package token

import "strconv"

// Token is the set of lexical tokens.
type Token int

// The lsit of all tokens
const (
	ILLEGAL Token = iota
	EOF

	// Identifiers and basic type literals
	INT
	STRING

	// Keywords
	CREATE_PARKING_LOT
	PARK
	LEAVE
	STATUS
	REGISTRATION_NUMBERS_FOR_CARS_WITH_COLOUR
	SLOT_NUMBERS_FOR_CARS_WITH_COLOUR
	SLOT_NUMBER_FOR_REGISTRATION_NUMBER
)

func (tok Token) String() string {
	s := ""
	if 0 <= tok && tok < Token(len(tokens)) {
		s = tokens[tok]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(tok)) + ")"
	}
	return s
}

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	INT:    "INT",
	STRING: "STRING",

	CREATE_PARKING_LOT: "create_parking_lot",
	PARK:               "park",
	LEAVE:              "leave",
	STATUS:             "status",
	REGISTRATION_NUMBERS_FOR_CARS_WITH_COLOUR: "registration_numbers_for_cars_with_colour",
	SLOT_NUMBERS_FOR_CARS_WITH_COLOUR:         "slot_numbers_for_cars_with_colour",
	SLOT_NUMBER_FOR_REGISTRATION_NUMBER:       "slot_number_for_registration_number",
}

var keywords = map[string]Token{
	"create_parking_lot": CREATE_PARKING_LOT,
	"park":               PARK,
	"leave":              LEAVE,
	"status":             STATUS,
	"registration_numbers_for_cars_with_colour": REGISTRATION_NUMBERS_FOR_CARS_WITH_COLOUR,
	"slot_numbers_for_cars_with_colour":         SLOT_NUMBERS_FOR_CARS_WITH_COLOUR,
	"slot_number_for_registration_number":       SLOT_NUMBER_FOR_REGISTRATION_NUMBER,
}

// Lookup maps an identifier to its keyword token or ILLEGAL (if not a keyword).
func Lookup(ident string) Token {
	if tok, isKeyword := keywords[ident]; isKeyword {
		return tok
	}
	return STRING
}
