package parking

import (
	"errors"
)

type ParkingLotService interface {
	Park(regNo, colour string) (*Slot, error)
	Leave(regNo string) (*Slot, error)
	LeaveByPosition(position int) (*Slot, error)
	FindByRegistrationNumber(numberPlate string) (*Slot, error)
	FindAllByColor(colour string) ([]*Slot, error)
}

type ParkingLot struct {
	MaxSlot int
	Slots   *Slot
}

func NewParkingLot(maxSlots int) *ParkingLot {
	return &ParkingLot{
		MaxSlot: maxSlots,
	}
}

func (p *ParkingLot) occupancy() int {
	var count int
	var s = p.Slots
	for s != nil {
		count++
		s = s.NextSlot
	}
	return count
}

func (p *ParkingLot) Park(regNo, colour string) (*Slot, error) {
	if p.occupancy() == p.MaxSlot {
		return nil, errors.New(MaxSlotReached)
	}

	car := NewCar(regNo, colour)

	if p.occupancy() == 0 {
		slot := NewSlot(car, 1)
		p.Slots = slot
		return slot, nil
	}

	if p.Slots.Position() > 1 {
		currSlot := p.Slots
		p.Slots = NewSlot(car, 1)
		p.Slots.AddNext(currSlot)
		currSlot.PrevSlot = p.Slots
	}

	slot := NewSlot(car, 0)
	p.Slots.AddNext(slot)

	return slot, nil
}

func (p *ParkingLot) Leave(regNo string) (*Slot, error) {
	if p.Slots == nil {
		return nil, errors.New(NoCarsParked)
	}

	slotFound, err := p.Slots.FindCar(regNo)
	if err != nil {
		return nil, errors.New(CarNotFound)
	}

	slotFound.Leave()
	if slotFound.PrevSlot == nil {
		p.Slots = slotFound.NextSlot
	}

	return slotFound, nil
}

func (p *ParkingLot) LeaveByPosition(position int) (*Slot, error) {
	if p.Slots == nil {
		return nil, errors.New(NoCarsParked)
	}

	slotFound, err := p.Slots.FindPosition(position)
	if err != nil {
		return nil, errors.New(CarNotFound)
	}

	slotFound.Leave()
	if slotFound.PrevSlot == nil {
		p.Slots = slotFound.NextSlot
	}

	return slotFound, nil
}

func (p *ParkingLot) FindByRegistrationNumber(numberPlate string) (*Slot, error) {
	if p.Slots == nil {
		return nil, errors.New(NoCarsParked)
	}

	return p.Slots.FindCar(numberPlate)
}

func (p *ParkingLot) FindAllByColor(colour string) ([]*Slot, error) {
	slotsList := []*Slot{}
	if p.Slots == nil {
		return nil, errors.New(NoCarsParked)
	}

	slots, _ := p.Slots.FindColor(colour)
	if len(slots) == 0 {
		return nil, errors.New(CarWithColorNotFound)
	}

	for i := range slots {
		slotsList = append(slotsList, slots[i])
	}
	return slotsList, nil
}

func (p *ParkingLot) AllSlots() ([]*Slot, error) {
	if p.Slots == nil {
		return nil, errors.New(NoCarsParked)
	}

	return p.Slots.ListSelf(), nil
}
