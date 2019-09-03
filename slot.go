package parking

import "errors"

type Slot struct {
	Car      *Car
	SlotNo   int
	PrevSlot *Slot
	NextSlot *Slot
}

func NewSlot(car *Car, slotNo int) *Slot {
	return &Slot{
		Car:    car,
		SlotNo: slotNo,
	}
}

func (s *Slot) Position() int {
	return s.SlotNo
}

func (s *Slot) AddNext(cs *Slot) error {
	if s.NextSlot == nil {
		s.NextSlot = cs.UpdatePosition(s.SlotNo + 1)
		cs.PrevSlot = s
		return nil
	}

	if s.NextSlot.SlotNo > (s.SlotNo + 1) {
		currentNext := s.NextSlot
		s.NextSlot = cs.UpdatePosition(s.SlotNo + 1)
		cs.PrevSlot = s
		cs.NextSlot = currentNext
		currentNext.PrevSlot = cs
		return nil
	}

	s.NextSlot.AddNext(cs)

	return nil
}

func (s *Slot) UpdatePosition(position int) *Slot {
	s.SlotNo = position
	return s
}

func (s *Slot) FindCar(regNo string) (*Slot, error) {
	if s.Car.RegNo == regNo {
		return s, nil
	}

	if s.NextSlot == nil {
		return nil, errors.New(CarNotFound)
	}

	return s.NextSlot.FindCar(regNo)
}

func (s *Slot) Leave() error {
	if s.PrevSlot != nil {
		s.PrevSlot.NextSlot = s.NextSlot
	}
	return nil
}

func (s *Slot) FindPosition(position int) (*Slot, error) {
	if s.SlotNo == position {
		return s, nil
	}

	if s.NextSlot == nil {
		return nil, errors.New(CarNotFound)
	}

	return s.NextSlot.FindPosition(position)
}

func (s *Slot) FindColor(colour string) ([]*Slot, error) {
	if s.Car.Colour == colour {
		if s.NextSlot == nil {
			return []*Slot{
				s,
			}, nil
		}

		slots, err := s.NextSlot.FindColor(colour)
		if err == nil {
			slots = append([]*Slot{s}, slots...)
		}
		return slots, err
	}

	if s.NextSlot == nil {
		return nil, errors.New(CarNotFound)
	}

	return s.NextSlot.FindColor(colour)
}

func (s *Slot) RegistrationNumber() string {
	if s.Car == nil {
		return ""
	}

	return s.Car.RegNo
}

func (s *Slot) Color() string {
	if s.Car == nil {
		return ""
	}

	return s.Car.Colour
}

func (s *Slot) ListSelf() []*Slot {
	if s.NextSlot == nil {
		return []*Slot{s}
	}

	return append([]*Slot{s}, s.NextSlot.ListSelf()...)
}
