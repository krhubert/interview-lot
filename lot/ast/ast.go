package ast

import (
	"fmt"

	"parking_lot/lot/token"
)

// Node represents an AST node.
type Node interface {
	String() string
}

// Statement represents a statement.
type Statement interface {
	Node
	statementNode()
}

// Program is a top-level AST node of a program.
type Program struct {
	Statements []Statement
}

// CreateParkingLotStatement represents a create parking lot statemant.
type CreateParkingLotStatement struct {
	Token  token.Token
	Number int
}

func (s *CreateParkingLotStatement) String() string {
	return fmt.Sprintf("%s %d", s.Token, s.Number)
}

// ParkStatement represents a park statemant.
type ParkStatement struct {
	Token              token.Token
	RegistrationNumber string
	Color              string
}

func (s *ParkStatement) String() string {
	return fmt.Sprintf("%s %s %s", s.Token, s.RegistrationNumber, s.Color)
}

// LeaveStatement represents a leave statemant.
type LeaveStatement struct {
	Token  token.Token
	Number int
}

func (s *LeaveStatement) String() string {
	return fmt.Sprintf("%s %d", s.Token, s.Number)
}

// StatusStatement represents a status statement.
type StatusStatement struct {
	Token token.Token
}

func (s *StatusStatement) String() string {
	return fmt.Sprintf("%s", s.Token)
}

// RegistrationNumbersForCarsWithColourStatement represents a registration statemant.
type RegistrationNumbersForCarsWithColourStatement struct {
	Token token.Token
	Color string
}

func (s *RegistrationNumbersForCarsWithColourStatement) String() string {
	return fmt.Sprintf("%s %s", s.Token, s.Color)
}

// SlotNumbersForCarsWithColourStatement represents a slot by color statement.
type SlotNumbersForCarsWithColourStatement struct {
	Token token.Token
	Color string
}

func (s *SlotNumbersForCarsWithColourStatement) String() string {
	return fmt.Sprintf("%s %s", s.Token, s.Color)
}

// SlotNumberForRegistrationNumberStatement represents a slot by number statement.
type SlotNumberForRegistrationNumberStatement struct {
	Token              token.Token
	RegistrationNumber string
}

func (s *SlotNumberForRegistrationNumberStatement) String() string {
	return fmt.Sprintf("%s %s", s.Token, s.RegistrationNumber)
}

// statementNode() ensures that only statement nodes can be assigned to a Statement.
func (*CreateParkingLotStatement) statementNode()                     {}
func (*ParkStatement) statementNode()                                 {}
func (*LeaveStatement) statementNode()                                {}
func (*StatusStatement) statementNode()                               {}
func (*RegistrationNumbersForCarsWithColourStatement) statementNode() {}
func (*SlotNumbersForCarsWithColourStatement) statementNode()         {}
func (*SlotNumberForRegistrationNumberStatement) statementNode()      {}
