package exec

import (
	"bytes"
	"testing"

	"parking_lot/database"
	"parking_lot/lot/ast"
)

var testPrograms = []struct {
	name    string
	program *ast.Program
}{
	{
		"empty program",
		&ast.Program{},
	},
	{
		"create car parking",
		&ast.Program{
			Statements: []ast.Statement{
				&ast.CreateParkingLotStatement{
					Number: 1,
				},
			},
		},
	},
	{
		"park car",
		&ast.Program{
			Statements: []ast.Statement{
				&ast.CreateParkingLotStatement{
					Number: 1,
				},
				&ast.ParkStatement{
					RegistrationNumber: "AA-00-AA-0000",
					Color:              "White",
				},
			},
		},
	},
	{
		"leave parked car",
		&ast.Program{
			Statements: []ast.Statement{
				&ast.CreateParkingLotStatement{
					Number: 1,
				},
				&ast.ParkStatement{
					RegistrationNumber: "AA-00-AA-0000",
					Color:              "White",
				},
				&ast.LeaveStatement{
					Number: 1,
				},
			},
		},
	},
}

func TestExecuteOutput(t *testing.T) {
	tests := []struct {
		stdout string
		stderr string
	}{
		{
			"",
			"",
		},
		{
			"Created a parking lot with 1 slots\n",
			"",
		},
		{
			"Created a parking lot with 1 slots\n" +
				"Allocated slot number: 1\n",
			"",
		},
		{
			"Created a parking lot with 1 slots\n" +
				"Allocated slot number: 1\n" +
				"Slot number 1 is free\n",
			"",
		},
	}

	var (
		stdout = new(bytes.Buffer)
		stderr = new(bytes.Buffer)
		e      = Executor{Stdout: stdout, Stderr: stderr}
		db     = database.NewDatabase(database.NewMemoryWriter())
	)

	for i, tt := range tests {
		e.Execute(testPrograms[i].program, db)
		if tt.stderr != stderr.String() {
			t.Errorf("test %s invalid stderr:\n\twant: %q\n\t got: %q", testPrograms[i].name, tt.stderr, stderr.String())
		}
		if tt.stdout != stdout.String() {
			t.Errorf("test %s invalid stderr:\n\twant: %q\n\t got: %q", testPrograms[i].name, tt.stdout, stdout.String())
		}
		stdout.Reset()
		stderr.Reset()
		db.Init(0)
	}
}

func TestExecuteDatabase(t *testing.T) {
	tests := []struct {
		cars []*database.Car
	}{
		{nil},
		{[]*database.Car{nil}},
		{[]*database.Car{
			database.MustNewCar("AA-00-AA-0000", "White"),
		}},
		{[]*database.Car{nil}},
	}

	var (
		stdout = new(bytes.Buffer)
		stderr = new(bytes.Buffer)
		e      = Executor{Stdout: stdout, Stderr: stderr}
	)

	for i, tt := range tests {
		db := database.NewDatabase(database.NewMemoryWriter())
		e.Execute(testPrograms[i].program, db)
		cars, _ := db.GetAll()
		if len(cars) != len(tt.cars) {
			t.Errorf("test %s invalid cars length - want: %d got: %c", testPrograms[i].name, len(tt.cars), len(cars))
		}

		for i := range cars {
			if tt.cars[i] == nil && cars[i] == nil {
				continue
			}

			if tt.cars[i].Color() != cars[i].Color() ||
				tt.cars[i].RegistrationNumber() != cars[i].RegistrationNumber() {
				t.Errorf("test %s invalid cars at position %d - want: %s got: %s", testPrograms[i].name, i, tt.cars[i], cars[i])
			}
		}
	}
}
