package database

import "testing"

func TestDatabaseFilterCars(t *testing.T) {
	db := NewDatabase(NewMemoryWriter())
	db.Init(2)
	db.Save(testCars[0])
	db.Save(testCars[1])

	cars, _ := db.FilterCars(FilterByRegistrationNumber("AA-00-A-000"))
	if len(cars) != 1 || cars[0].registrationNumber != "AA-00-A-000" {
		t.Errorf("filter by registration number failed")
	}

	cars, _ = db.FilterCars(FilterByColor("White"))
	if len(cars) != 1 || cars[0].color != "White" {
		t.Errorf("filter by color failed")
	}

	cars, _ = db.FilterCars(nil)
	if len(cars) != 2 {
		t.Errorf("nil filter failed")
	}
}

func TestDatabaseFilterSlotNumbers(t *testing.T) {
	db := NewDatabase(NewMemoryWriter())
	db.Init(2)
	db.Save(testCars[0])
	db.Save(testCars[1])

	slots, _ := db.FilterSlotNumbers(FilterByRegistrationNumber("AA-00-A-001"))
	if len(slots) != 1 || slots[0] != 1 {
		t.Errorf("filter by registration number failed")
	}

	slots, _ = db.FilterSlotNumbers(FilterByColor("Black"))
	if len(slots) != 1 || slots[0] != 1 {
		t.Errorf("filter by color failed")
	}

	slots, _ = db.FilterSlotNumbers(nil)
	if len(slots) != 2 {
		t.Errorf("filter by color failed")
	}
}
