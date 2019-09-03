package parking

import (
	"errors"
)

type ParkingLotService interface {
	Park(regNo, colour string) (*slot, error)
	Leave(regNo string) (*slot, error)
	LeaveByPosition(position int) (*slot, error)
	FindByRegistrationNumber(numberPlate string) (*slot, error)
	FindAllByColor(colour string) ([]*slot, error)
}

// parkingLot stores the slots for a parking lot and the max slots available in it
type parkingLot struct {
	maxSlot int
	slots   *slot
}

// NewParkingLot creates new parking lot with the a specific number of slots
func NewParkingLot(maxSlots int) *parkingLot {
	return &parkingLot{
		maxSlot: maxSlots,
	}
}

// occupancy returns the number of slots occupied
func (p *parkingLot) occupancy() int {
	var count int
	var s = p.slots
	for s != nil {
		count++
		s = s.nextSlot
	}
	return count
}

// Park, parks a car into the next avaliable slot.
func (p *parkingLot) Park(regNo, colour string) (*slot, error) {
	if p.occupancy() == p.maxSlot {
		return nil, errors.New(MaxSlotReached)
	}

	car := NewCar(regNo, colour)

	if p.occupancy() == 0 {
		slot := NewSlot(car, 1)
		p.slots = slot
		return slot, nil
	}

	if p.slots.Position() > 1 {
		currSlot := p.slots
		p.slots = NewSlot(car, 1)
		p.slots.AddNext(currSlot)
		currSlot.prevSlot = p.slots
	}

	slot := NewSlot(car, 0)
	p.slots.AddNext(slot)

	return slot, nil
}

// Leave deallocates a specific regNo car from its slot
func (p *parkingLot) Leave(regNo string) (*slot, error) {
	if p.slots == nil {
		return nil, errors.New(NoCarsParked)
	}

	slotFound, err := p.slots.FindCar(regNo)
	if err != nil {
		return nil, errors.New(CarNotFound)
	}

	slotFound.Leave()
	if slotFound.prevSlot == nil {
		p.slots = slotFound.nextSlot
	}

	return slotFound, nil
}

// LeaveByPosition deallocates a car from the specific position if it is parked
func (p *parkingLot) LeaveByPosition(position int) (*slot, error) {
	if p.slots == nil {
		return nil, errors.New(NoCarsParked)
	}

	slotFound, err := p.slots.FindPosition(position)
	if err != nil {
		return nil, errors.New(CarNotFound)
	}

	slotFound.Leave()
	if slotFound.prevSlot == nil {
		p.slots = slotFound.nextSlot
	}

	return slotFound, nil
}

// FindByRegistrationNumber finds a specific car with the RegNo
func (p *parkingLot) FindByRegistrationNumber(numberPlate string) (*slot, error) {
	if p.slots == nil {
		return nil, errors.New(NoCarsParked)
	}

	return p.slots.FindCar(numberPlate)
}

// FindAllByColor, finds all the cars with the particular colour
func (p *parkingLot) FindAllByColor(colour string) ([]*slot, error) {
	slotsList := []*slot{}
	if p.slots == nil {
		return nil, errors.New(NoCarsParked)
	}

	slots, _ := p.slots.FindColour(colour)
	if len(slots) == 0 {
		return nil, errors.New(CarWithColorNotFound)
	}

	for i := range slots {
		slotsList = append(slotsList, slots[i])
	}
	return slotsList, nil
}

// AllSlots returns the list of all occupied slots
func (p *parkingLot) AllSlots() ([]*slot, error) {
	if p.slots == nil {
		return nil, errors.New(NoCarsParked)
	}

	return p.slots.ListSelf(), nil
}
