package database

import (
	"errors"
	"fmt"
)

// Writer is an interface for stroing the cars.
type Writer interface {
	Init(capacity int) error
	Save(car *Car) (int, error)
	Remove(pos int) error
	GetAll() ([]*Car, error)
}

// ErrOutOfRange is out of range error.
type ErrOutOfRange struct {
	pos      int
	capacity int
}

func (e *ErrOutOfRange) Error() string {
	return fmt.Sprintf("slot number %d out of range [1, %d]", e.pos, e.capacity)
}

var (
	// ErrFull is returned when there is no more parking slots.
	ErrFull = errors.New("sorry, parking lot is full")
	// ErrIdentity is returned when two cars with the same registration number
	// are saved.
	ErrIdentity = errors.New("identity thieves are not welcome, calling police")
)

// MemoryWriter is writer that keeps everything in memory.
type MemoryWriter struct {
	cars []*Car
}

// NewMemoryWriter creates new memory writer.
func NewMemoryWriter() *MemoryWriter {
	return &MemoryWriter{}
}

// Init initializes writer with given capacity.
// Call Init again will remove all cars from current writer.
func (w *MemoryWriter) Init(capacity int) error {
	w.cars = make([]*Car, capacity)
	return nil
}

// Save saves given car in the first free slot.
func (w *MemoryWriter) Save(car *Car) (int, error) {
	saveIndex := -1
	for i := range w.cars {
		if w.cars[i] != nil && w.cars[i].registrationNumber == car.registrationNumber {
			return -1, ErrIdentity
		}

		if w.cars[i] == nil && saveIndex == -1 {
			saveIndex = i
			break
		}
	}

	if saveIndex == -1 {
		return -1, ErrFull
	}

	w.cars[saveIndex] = car
	return saveIndex, nil
}

// Remove removes cars from given position.
func (w *MemoryWriter) Remove(pos int) error {
	if pos < 0 || pos >= cap(w.cars) {
		return &ErrOutOfRange{pos, cap(w.cars)}
	}
	w.cars[pos] = nil
	return nil
}

// GetAll returns all the cars.
func (w *MemoryWriter) GetAll() ([]*Car, error) {
	return w.cars, nil
}

// FileWriter writes cars info to file.
type FileWriter struct{}

// NewFileWriter creates new file writer.
func NewFileWriter(file string) (*FileWriter, error) {
	return nil, errors.New("unimplemented")
}

// Init initializes writer with given capacity.
func (w *FileWriter) Init(capacity int) error {
	return errors.New("unimplemented")
}

// Save saves given car in the first free slot.
func (w *FileWriter) Save(car *Car) (int, error) {
	return -1, errors.New("unimplemented")
}

// Remove removes cars from given position.
func (w *FileWriter) Remove(pos int) error {
	return errors.New("unimplemented")
}

// GetAll returns all the cars.
func (w *FileWriter) GetAll() ([]*Car, error) {
	return nil, errors.New("unimplemented")
}
