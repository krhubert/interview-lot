package database

import "testing"

var (
	testCars = []*Car{
		{
			registrationNumber: "AA-00-A-000",
			color:              "White",
		},
		{
			registrationNumber: "AA-00-A-001",
			color:              "Black",
		},
	}

	// car used to test if parking is full
	extraTestCar = &Car{
		registrationNumber: "AA-00-A-002",
		color:              "Red",
	}
)

func TestMemoryWriterInit(t *testing.T) {
	w := NewMemoryWriter()

	w.Init(2)
	if c := cap(w.cars); c != 2 {
		t.Fatalf("invalid capacity - want: %d, got: %d ", 2, c)
	}

	w.Init(0)
	if l := cap(w.cars); l != 0 {
		t.Fatalf("reinit invalid capacity - want: %d, got: %d", 0, l)
	}
}

func TestMemoryWriterSave(t *testing.T) {
	w := NewMemoryWriter()
	w.Init(2)
	si, err := w.Save(testCars[0])
	if err != nil {
		t.Fatalf("save error: %s", err)
	}
	if si != 0 {
		t.Fatalf("save on invalid slot - want: %d, got: %d", 0, si)
	}

	si, err = w.Save(testCars[1])
	if err != nil {
		t.Fatalf("save error: %s", err)
	}
	if si != 1 {
		t.Fatalf("save on invalid slot - want: %d, got: %d", 1, si)
	}

	// test parking is full
	if _, err := w.Save(extraTestCar); err != ErrFull {
		t.Fatalf("save error - want: %s, got: %s", ErrFull, err)
	}

	// test save with the same registration number
	if _, err := w.Save(testCars[0]); err != ErrIdentity {
		t.Fatalf("save error - want: %s, got: %s", ErrFull, err)
	}

	for i := 0; i < len(testCars); i++ {
		if w.cars[i].registrationNumber != testCars[i].registrationNumber {
			t.Fatalf("car[%d] has invalid registration number - want: %s, got: %s",
				i, testCars[i].registrationNumber, w.cars[i].registrationNumber)
		}
		if w.cars[i].color != testCars[i].color {
			t.Fatalf("car[%d] has invalid color - want: %s, got: %s",
				i, testCars[i].color, w.cars[i].color)
		}
	}
}

func TestMemoryWriterRemove(t *testing.T) {
	w := NewMemoryWriter()

	w.Init(len(testCars))
	for _, car := range testCars {
		w.Save(car)
	}

	if err := w.Remove(0); err != nil {
		t.Fatalf("remove unexpected error: %s", err)
	}

	if w.cars[0] != nil {
		t.Fatalf("car not removed")
	}

	// test remove errors
	if err, ok := w.Remove(-1).(*ErrOutOfRange); !ok {
		t.Fatalf("remove expected ErrOutOfRange but got: %s", err)
	}
	if err, ok := w.Remove(len(testCars)).(*ErrOutOfRange); !ok {
		t.Fatalf("remove expected ErrOutOfRange but got: %s", err)
	}
}

func TestMemoryWriterGetAll(t *testing.T) {
	w := NewMemoryWriter()

	w.Init(len(testCars))
	for _, car := range testCars {
		w.Save(car)
	}

	newCars, _ := w.GetAll()
	if len(newCars) != len(testCars) {
		t.Fatalf("get all returned invalid number of cars - want: %d, got: %d", len(testCars), len(newCars))
	}

	w.Remove(1)
	newCars, _ = w.GetAll()
	if len(newCars) != len(testCars) {
		t.Fatalf("get all returned invalid number of cars - want: %d, got: %d", len(testCars), len(newCars))
	}
}
