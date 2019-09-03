package parking

import (
	"fmt"
	"strconv"
	"strings"
)

type resp struct {
	parkingLot *parkingLot
	slots      []*slot
	command    Input
}

func (r resp) String() string {
	switch r.command {
	case CreateParking:
		return fmt.Sprintf("Created a parking lot with %d slots", r.parkingLot.maxSlot)
	case Park:
		return fmt.Sprintf("Allocated slot number: %d", r.slots[0].Position())
	case Status:
		content := fmt.Sprintf("Slot No.\tRegistration No\tColor")
		for _, slot := range r.slots {
			content += fmt.Sprintf("\n%d %s %s", slot.Position(), slot.RegistrationNumber(), slot.Colour())
		}
		return content
	case Leave:
		return fmt.Sprintf("Slot number %d is free", r.slots[0].Position())
	case CarRegNoWithdColour:
		regNumbers := []string{}
		for _, s := range r.slots {
			regNumbers = append(regNumbers, s.car.RegNo)
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
	case NotFisrt:
		return fmt.Sprintf("parking already created")
	default:
		return fmt.Sprintf("Invalid input")
	}

	return ""
}

// NewResponse return an DbResponse Object.
func NewResponse(pl *parkingLot) *resp {
	return &resp{pl, nil, CreateParking}
}
