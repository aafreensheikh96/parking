package parking

import (
	"fmt"
	"strconv"
	"strings"
)

type Parking struct {
	parkingLot *ParkingLot
}

func NewParking(maxSlots int) *Parking {
	pl := NewParkingLot(maxSlots)
	return &Parking{
		pl,
	}
}

type Resp struct {
	parkingLot *ParkingLot
	slots      []*Slot
	command    Input
}

func (r Resp) String() string {
	switch r.command {
	case CreateParking:
		return fmt.Sprintf("Created a parking lot with %d slots", r.parkingLot.MaxSlot)
	case Park:
		return fmt.Sprintf("Allocated slot number: %d", r.slots[0].Position())
	case Status:
		content := fmt.Sprintf("Slot No.\tRegistration No\tColor")
		for _, slot := range r.slots {
			content += fmt.Sprintf("\n%d %s %s", slot.Position(), slot.RegistrationNumber(), slot.Color())
		}
		return content
	case Leave:
		return fmt.Sprintf("Slot number %d is free", r.slots[0].Position())
	case CarRegNoWithdColour:
		regNumbers := []string{}
		for _, s := range r.slots {
			regNumbers = append(regNumbers, s.Car.RegNo)
		}
		return strings.Join(regNumbers, ", ")
	case SlotWithColour:
		positions := []string{}
		for _, s := range r.slots {
			positions = append(positions, strconv.Itoa(s.Position()))
		}

		return strings.Join(positions, ", ")
	case SlotWithRegNo:
		return strconv.Itoa(r.slots[0].Position())
	}
	return ""
}

// NewResponse return an DbResponse Object.
func NewResponse(pl *ParkingLot) *Resp {
	return &Resp{pl, nil, CreateParking}
}

func (p *Parking) Park(numberPlate, color string) (*Resp, error) {
	slot, err := p.parkingLot.Park(numberPlate, color)
	sResponse := &Resp{
		slots: []*Slot{
			slot,
		},
		command: Park,
	}

	return sResponse, err
}

func (p *Parking) LeaveByPosition(position int) (*Resp, error) {
	slot, err := p.parkingLot.LeaveByPosition(position)
	sResponse := &Resp{
		slots: []*Slot{
			slot,
		},
		command: Leave,
	}

	return sResponse, err
}

func (p *Parking) FindByRegistrationNumber(numberPlate string) (*Resp, error) {
	slot, err := p.parkingLot.FindByRegistrationNumber(numberPlate)
	sResponse := &Resp{
		slots: []*Slot{
			slot,
		},
		command: SlotWithRegNo,
	}

	return sResponse, err
}

func (p *Parking) FindAllByColor(color string) (*Resp, error) {
	slots, err := p.parkingLot.FindAllByColor(color)
	sResponse := &Resp{
		slots:   slots,
		command: SlotWithColour,
	}

	return sResponse, err
}

func (p *Parking) All() (*Resp, error) {
	slots, err := p.parkingLot.AllSlots()
	return &Resp{
		slots:   slots,
		command: Status,
	}, err
}
