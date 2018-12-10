package database

import (
	"fmt"
	"regexp"
)

// Colors is list of all valid colors.
var Colors = []string{
	"White",
	"Yellow",
	"Red",
	"Blue",
	"Green",
	"Brown",
	"Black",
}

// regexp to validate registration number.
var validRegistrationNumber = regexp.MustCompile(`^[A-Z]{2}-\d{2}-[A-Z]{1,2}-\d{3,4}$`)

// Car holds information like color and registration number.
type Car struct {
	registrationNumber string
	color              string
}

// NewCar creates a car.
func NewCar(registrationNumber, color string) (*Car, error) {
	if !validRegistrationNumber.Match([]byte(registrationNumber)) {
		return nil, fmt.Errorf("car registration number %q is invalid", registrationNumber)
	}

	found := false
	for i := range Colors {
		if Colors[i] == color {
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("car colour %q is invalid", color)
	}

	return &Car{
		registrationNumber: registrationNumber,
		color:              color,
	}, nil
}

// MustNewCar creates a car. It panics on error.
func MustNewCar(registrationNumber, color string) *Car {
	c, err := NewCar(registrationNumber, color)
	if err != nil {
		panic(err)
	}
	return c
}

// RegistrationNumber returns car regitration number.
func (c *Car) RegistrationNumber() string {
	return c.registrationNumber
}

// Color returns car color.
func (c *Car) Color() string {
	return c.color
}

func (c *Car) String() string {
	return c.registrationNumber + " " + c.color
}
