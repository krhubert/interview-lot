package scanner

import (
	"parking_lot/lot/token"
)

// end of file
const eof byte = 0

// A Scanner holds the scanner's internal state while processing a lot source.
type Scanner struct {
	src string

	ch         byte // current character
	offset     int  // character offset
	readOffset int  // reading offset (position after current character)
}

// New returns new scanner.
func New(src string) *Scanner {
	s := &Scanner{src: src}
	s.next()
	return s
}

func (s *Scanner) next() {
	if s.readOffset >= len(s.src) {
		s.ch = eof
	} else {
		s.ch = s.src[s.readOffset]
	}
	s.offset = s.readOffset
	s.readOffset++
}

func (s *Scanner) skipWhitespace() {
	for isWhitespace(s.ch) {
		s.next()
	}
}

func (s *Scanner) scanString() string {
	offset := s.offset
	for !isWhitespace(s.ch) && s.ch != eof {
		s.next()
	}
	return s.src[offset:s.offset]
}

func (s *Scanner) scanNumber() string {
	offset := s.offset
	for isDigit(s.ch) {
		s.next()
	}
	return s.src[offset:s.offset]
}

// Scan scans the next token and returns the token position, the token,
// and its literal string if applicable. The source end is indicated by
// token.EOF.
func (s *Scanner) Scan() (pos int, tok token.Token, lit string) {
	s.skipWhitespace()
	pos = s.offset

	switch {
	case isLetter(s.ch):
		lit = s.scanString()
		tok = token.Lookup(lit)
	case isDigit(s.ch):
		lit = s.scanNumber()
		tok = token.INT
	case s.ch == eof:
		tok = token.EOF
	default:
		tok = token.ILLEGAL
		lit = string(s.ch)
	}
	return pos, tok, lit
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
