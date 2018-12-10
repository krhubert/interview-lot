package exec

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"parking_lot/database"
	"parking_lot/lot/ast"
)

// Executor handles program execution and output.
type Executor struct {
	Stdout io.Writer
	Stderr io.Writer
}

// NewExecutor createx new Executor with stdout and stderr set to os.Stdout
// NOTE: IMPORTANT: Stderr is set to os.Stdout.
func NewExecutor() *Executor {
	return &Executor{
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}
}

// Execute executes lot program on given database.
func (e *Executor) Execute(program *ast.Program, db *database.Database) {
	for _, stmt := range program.Statements {
		switch stmt := stmt.(type) {
		case *ast.CreateParkingLotStatement:
			e.execCreateParkingLotStatement(db, stmt)
		case *ast.ParkStatement:
			e.execParkStatement(db, stmt)
		case *ast.LeaveStatement:
			e.execLeaveStatement(db, stmt)
		case *ast.StatusStatement:
			e.execStatusStatement(db)
		case *ast.RegistrationNumbersForCarsWithColourStatement:
			e.execRegistrationNumbersForCarsWithColourStatement(db, stmt)
		case *ast.SlotNumbersForCarsWithColourStatement:
			e.execSlotNumbersForCarsWithColourStatement(db, stmt)
		case *ast.SlotNumberForRegistrationNumberStatement:
			e.execSlotNumberForRegistrationNumberStatement(db, stmt)
		}
	}
}

func (e *Executor) execCreateParkingLotStatement(db *database.Database, stmt *ast.CreateParkingLotStatement) {
	if err := db.Init(stmt.Number); err != nil {
		fmt.Fprintln(e.Stderr, err)
	} else {
		fmt.Fprintf(e.Stdout, "Created a parking lot with %d slots\n", stmt.Number)
	}
}

func (e *Executor) execParkStatement(db *database.Database, stmt *ast.ParkStatement) {
	car, err := database.NewCar(stmt.RegistrationNumber, stmt.Color)
	if err != nil {
		fmt.Fprintln(e.Stderr, err)
		return
	}

	if i, err := db.Save(car); err != nil {
		fmt.Fprintln(e.Stderr, err)
	} else {
		fmt.Fprintf(e.Stdout, "Allocated slot number: %d\n", i+1)
	}
}

func (e *Executor) execLeaveStatement(db *database.Database, stmt *ast.LeaveStatement) {
	if err := db.Remove(stmt.Number - 1); err != nil {
		fmt.Fprintln(e.Stderr, err)
	} else {
		fmt.Fprintf(e.Stdout, "Slot number %d is free\n", stmt.Number)
	}
}

func (e *Executor) execStatusStatement(db *database.Database) {
	cars, err := db.GetAll()
	if err != nil {
		fmt.Fprintln(e.Stderr, err)
		return
	}

	w := tabwriter.NewWriter(e.Stdout, 0, 0, 4, ' ', 0)
	fmt.Fprintf(w, "Slot No.\tRegistration No\tColour\n")
	for i, car := range cars {
		if car != nil {
			fmt.Fprintf(w, "%d\t%s\t%s\n", i+1, car.RegistrationNumber(), car.Color())
		}
	}
	if err := w.Flush(); err != nil {
		fmt.Fprintln(e.Stderr, err)
	}
}

func (e *Executor) execRegistrationNumbersForCarsWithColourStatement(db *database.Database, stmt *ast.RegistrationNumbersForCarsWithColourStatement) {
	cars, err := db.FilterCars(database.FilterByColor(stmt.Color))
	if err != nil {
		fmt.Fprintln(e.Stderr, err)
		return
	}
	if len(cars) == 0 {
		fmt.Fprintf(e.Stderr, "Not found\n")
	} else {
		var s []string
		for _, car := range cars {
			s = append(s, car.RegistrationNumber())
		}
		fmt.Fprintln(e.Stdout, strings.Join(s, ", "))
	}
}

func (e *Executor) execSlotNumbersForCarsWithColourStatement(db *database.Database, stmt *ast.SlotNumbersForCarsWithColourStatement) {
	slots, err := db.FilterSlotNumbers(database.FilterByColor(stmt.Color))
	if err != nil {
		fmt.Fprintln(e.Stderr, err)
		return
	}

	if len(slots) == 0 {
		fmt.Fprintf(e.Stderr, "Not found\n")
	} else {
		fmt.Fprintln(e.Stdout, intSliceToString(slots))
	}
}

func (e *Executor) execSlotNumberForRegistrationNumberStatement(db *database.Database, stmt *ast.SlotNumberForRegistrationNumberStatement) {
	slots, err := db.FilterSlotNumbers(database.FilterByRegistrationNumber(stmt.RegistrationNumber))
	if err != nil {
		fmt.Fprintln(e.Stderr, err)
		return
	}

	if len(slots) == 0 {
		fmt.Fprintf(e.Stderr, "Not found\n")
	} else {
		fmt.Fprintln(e.Stdout, intSliceToString(slots))
	}
}

// intSliceToString join given int slice into string.
// It uses ", " as separator.
// IMPORATANT: it adds +1 to every int to keep program output consistent
// with specification.
func intSliceToString(a []int) string {
	if len(a) == 0 {
		return ""
	}

	s := strconv.FormatInt(int64(a[0]+1), 10)
	for i := 1; i < len(a); i++ {
		s += ", " + strconv.FormatInt(int64(a[i]+1), 10)
	}
	return s
}
