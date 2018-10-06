package database

// Database handles filter and searching of cars.
type Database struct {
	Writer
}

// NewDatabase creates new database with given writer.
func NewDatabase(w Writer) *Database {
	return &Database{Writer: w}
}

// FilterCars filters cars with given filter and returns them.
// Passing nil fillter will cause in returing all cars.
func (db *Database) FilterCars(fok Filter) ([]*Car, error) {
	cars, err := db.GetAll()
	if err != nil {
		return nil, err
	}

	var newCars []*Car
	for _, car := range cars {
		if fok == nil || fok(car) {
			newCars = append(newCars, car)
		}
	}
	return newCars, nil
}

// FilterSlotNumbers filters cars with given filter and returns it's position in database.
// Passing nil fillter will cause in returing all cars slots.
func (db *Database) FilterSlotNumbers(fok Filter) ([]int, error) {
	cars, err := db.GetAll()
	if err != nil {
		return nil, err
	}

	var slots []int
	for i, car := range cars {
		if fok == nil || fok(car) {
			slots = append(slots, i)
		}
	}
	return slots, nil
}

// Filter is a function for filtering cars.
type Filter func(*Car) bool

// FilterByRegistrationNumber filters cars by registration number.
func FilterByRegistrationNumber(registrationNumber string) Filter {
	return func(c *Car) bool {
		return c.registrationNumber == registrationNumber
	}
}

// FilterByColor filters cars by color.
func FilterByColor(color string) Filter {
	return func(c *Car) bool {
		return c.color == color
	}
}
