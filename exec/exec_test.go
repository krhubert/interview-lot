package exec

import (
	"bytes"
	"testing"

	"parking_lot/database"
	"parking_lot/lot/ast"
)

func TestExecuteOutput(t *testing.T) {
	tests := []struct {
		name        string
		program     *ast.Program
		expectedOut string
		expectedErr string
	}{
		{
			"empty program",
			&ast.Program{},
			"",
			"",
		},
		{
			"car parking",
			&ast.Program{
				Statements: []ast.Statement{
					&ast.CreateParkingLotStatement{
						Number: 1,
					},
				},
			},
			"Created a parking lot with 1 slots\n",
			"",
		},
		// TODO: add test for each statement
	}

	var (
		stdout = new(bytes.Buffer)
		stderr = new(bytes.Buffer)
		e      = Executor{Stdout: stdout, Stderr: stderr}
	)

	for _, tt := range tests {
		db := database.NewDatabase(database.NewMemoryWriter())
		e.Execute(tt.program, db)
		if tt.expectedErr != stderr.String() {
			t.Errorf("test %s invalid stderr:\n\twant: %q\n\tgot: %q", tt.name, tt.expectedErr, stderr.String())
		}
		if tt.expectedOut != stdout.String() {
			t.Errorf("test %s invalid stderr:\n\twant: %q\n\tgot: %q", tt.name, tt.expectedOut, stdout.String())
		}
	}
}

func TestExecuteDatabase(t *testing.T) {
	tests := []struct {
		name    string
		program *ast.Program
		cars    []*database.Car
	}{
		{
			"empty program",
			&ast.Program{},
			nil,
		},
		{
			"car parking",
			&ast.Program{
				Statements: []ast.Statement{
					&ast.CreateParkingLotStatement{
						Number: 1,
					},
				},
			},
			[]*database.Car{nil},
		},
		// TODO: add test for each statement
	}

	var (
		stdout = new(bytes.Buffer)
		stderr = new(bytes.Buffer)
		e      = Executor{Stdout: stdout, Stderr: stderr}
	)

	for _, tt := range tests {
		db := database.NewDatabase(database.NewMemoryWriter())
		e.Execute(tt.program, db)
		cars, _ := db.GetAll()
		if len(cars) != len(tt.cars) {
			t.Errorf("test %s invalid cars length - want: %d got: %c", tt.name, len(tt.cars), len(cars))
		}

		for i := range cars {
			if tt.cars[i] == nil && cars[i] == nil {
				continue
			}

			if tt.cars[i].Color() != cars[i].Color() ||
				tt.cars[i].RegistrationNumber() != cars[i].RegistrationNumber() {
				t.Errorf("test %s invalid cars at position %d - want: %s got: %s", tt.name, i, tt.cars[i], cars[i])
			}
		}
	}
}
