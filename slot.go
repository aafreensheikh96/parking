package parking

import (
	"errors"
)

type slot struct {
	car      *car
	slotNo   int
	prevSlot *slot
	nextSlot *slot
}

// NewSlot creates a slor for a car with specific regNo and colour
func NewSlot(car *car, slotNo int) *slot {
	return &slot{
		car:    car,
		slotNo: slotNo,
	}
}

// Position returns the slot no of a slot
func (s *slot) Position() int {
	return s.slotNo
}

// AddNext adds a slot next to the current slot
// in the parking lot if it is not completely occupied
func (s *slot) AddNext(cs *slot) error {
	if s.nextSlot == nil {
		s.nextSlot = cs.UpdatePosition(s.slotNo + 1)
		cs.prevSlot = s
		return nil
	}

	if s.nextSlot.slotNo > (s.slotNo + 1) {
		currentNext := s.nextSlot
		s.nextSlot = cs.UpdatePosition(s.slotNo + 1)
		cs.prevSlot = s
		cs.nextSlot = currentNext
		currentNext.prevSlot = cs
		return nil
	}

	s.nextSlot.AddNext(cs)

	return nil
}

// UpdatePosition updates the position of a slot
func (s *slot) UpdatePosition(position int) *slot {
	s.slotNo = position
	return s
}

// FindCar, finds a slot of a car with specific regNo
func (s *slot) FindCar(regNo string) (*slot, error) {
	if s.car.RegNo == regNo {
		return s, nil
	}

	if s.nextSlot == nil {
		return nil, errors.New(CarNotFound)
	}

	return s.nextSlot.FindCar(regNo)
}

// Leave deallocates a slot
func (s *slot) Leave() error {
	if s.prevSlot != nil {
		s.prevSlot.nextSlot = s.nextSlot
	}
	return nil
}

// FindPosition finds a slot with specific position
func (s *slot) FindPosition(position int) (*slot, error) {
	if s.slotNo == position {
		return s, nil
	}

	if s.nextSlot == nil {
		return nil, errors.New(CarNotFound)
	}

	return s.nextSlot.FindPosition(position)
}

// FindColour finds a car with specific colour
func (s *slot) FindColour(colour string) ([]*slot, error) {

	if s.car.Colour == colour {
		if s.nextSlot == nil {
			return []*slot{
				s,
			}, nil
		}
		slots, err := s.nextSlot.FindColour(colour)
		if err == nil {
			slots = append([]*slot{s}, slots...)
		}
		return slots, nil
	}

	if s.nextSlot == nil {
		return nil, nil
	}

	return s.nextSlot.FindColour(colour)
}

// RegistrationNumber returns the regNo of a car at the slot
func (s *slot) RegistrationNumber() string {
	if s.car == nil {
		return ""
	}

	return s.car.RegNo
}

// Colour returns the Colour of a car at the slot
func (s *slot) Colour() string {
	if s.car == nil {
		return ""
	}

	return s.car.Colour
}

// ListSelf lists the slot if it is not nil
func (s *slot) ListSelf() []*slot {
	if s.nextSlot == nil {
		return []*slot{s}
	}

	return append([]*slot{s}, s.nextSlot.ListSelf()...)
}
