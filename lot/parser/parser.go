package parser

import (
	"fmt"
	"strconv"

	"parking_lot/errors"
	"parking_lot/lot/ast"
	"parking_lot/lot/scanner"
	"parking_lot/lot/token"
)

// The parser structure holds the parser's internal state.
type parser struct {
	scanner *scanner.Scanner
	errors  []error

	// next token
	pos int         // token position
	tok token.Token // one token look-ahead
	lit string      // token literal
}

// newParser returns a new parser.
func newParser(src string) *parser {
	p := &parser{scanner: scanner.New(src)}
	p.next()
	return p
}

// next advance to the next token.
func (p *parser) next() {
	p.pos, p.tok, p.lit = p.scanner.Scan()
}

func (p *parser) expect(tok token.Token) bool {
	if p.next(); p.tok != tok {
		p.errors = append(p.errors, fmt.Errorf("unexpected token %q at pos %d, expecting %q", p.lit, p.pos, tok))
		return false
	}
	return true
}

func (p *parser) parseStatement() ast.Statement {
	switch p.tok {
	case token.CREATE_PARKING_LOT:
		if stmt := p.parseCreateParkingLot(); stmt != nil {
			return stmt
		}
	case token.PARK:
		if stmt := p.parsePark(); stmt != nil {
			return stmt
		}
	case token.LEAVE:
		if stmt := p.parseLeave(); stmt != nil {
			return stmt
		}
	case token.STATUS:
		if stmt := p.parseStatus(); stmt != nil {
			return stmt
		}
	case token.REGISTRATION_NUMBERS_FOR_CARS_WITH_COLOUR:
		if stmt := p.parseRegistrationNumbersForCarsWithColour(); stmt != nil {
			return stmt
		}
	case token.SLOT_NUMBERS_FOR_CARS_WITH_COLOUR:
		if stmt := p.parseSlotNumbersForarsWithColour(); stmt != nil {
			return stmt
		}
	case token.SLOT_NUMBER_FOR_REGISTRATION_NUMBER:
		if stmt := p.parseSlotNumberForRegistrationNumber(); stmt != nil {
			return stmt
		}
	default:
		p.errors = append(p.errors, fmt.Errorf("unexpected token %q at pos %d", p.lit, p.pos))
		return nil
	}
	return nil
}

func (p *parser) parseCreateParkingLot() *ast.CreateParkingLotStatement {
	if !p.expect(token.INT) {
		return nil
	}

	n, err := strconv.ParseInt(p.lit, 10, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("invalid number %q at pos %d", p.lit, p.pos))
		return nil
	}

	return &ast.CreateParkingLotStatement{
		Token:  token.CREATE_PARKING_LOT,
		Number: int(n),
	}
}

func (p *parser) parsePark() *ast.ParkStatement {
	if !p.expect(token.STRING) {
		return nil
	}
	registrationNumber := p.lit

	if !p.expect(token.STRING) {
		return nil
	}
	color := p.lit

	return &ast.ParkStatement{
		Token:              token.PARK,
		RegistrationNumber: registrationNumber,
		Color:              color,
	}
}

func (p *parser) parseLeave() *ast.LeaveStatement {
	if !p.expect(token.INT) {
		return nil
	}

	n, err := strconv.ParseInt(p.lit, 10, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("invalid number %q at pos %d", p.lit, p.pos))
		return nil
	}

	return &ast.LeaveStatement{
		Token:  token.CREATE_PARKING_LOT,
		Number: int(n),
	}
}

func (p *parser) parseStatus() *ast.StatusStatement {
	return &ast.StatusStatement{Token: token.STATUS}
}

func (p *parser) parseRegistrationNumbersForCarsWithColour() *ast.RegistrationNumbersForCarsWithColourStatement {
	if !p.expect(token.STRING) {
		return nil
	}
	color := p.lit

	return &ast.RegistrationNumbersForCarsWithColourStatement{
		Token: token.PARK,
		Color: color,
	}
}

func (p *parser) parseSlotNumbersForarsWithColour() *ast.SlotNumbersForCarsWithColourStatement {
	if !p.expect(token.STRING) {
		return nil
	}
	color := p.lit

	return &ast.SlotNumbersForCarsWithColourStatement{
		Token: token.PARK,
		Color: color,
	}
}

func (p *parser) parseSlotNumberForRegistrationNumber() *ast.SlotNumberForRegistrationNumberStatement {
	if !p.expect(token.STRING) {
		return nil
	}
	registrationNumber := p.lit

	return &ast.SlotNumberForRegistrationNumberStatement{
		Token:              token.PARK,
		RegistrationNumber: registrationNumber,
	}
}

// Parse parses the lot source code and returns a new Program AST node.
func Parse(src string) (*ast.Program, error) {
	program := &ast.Program{
		Statements: []ast.Statement{},
	}

	p := newParser(src)
	for p.tok != token.EOF {
		if stmt := p.parseStatement(); stmt != nil {
			program.Statements = append(program.Statements, stmt)
		} else {
			break
		}
		p.next()
	}

	if len(p.errors) > 0 {
		return nil, errors.Join(p.errors)
	}
	return program, nil
}
