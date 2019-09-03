package parking

import (
	"errors"
	"strconv"
	"strings"
)

func parseCommand(command string) []string {
	parsedCommand := []string{}

	command = strings.Replace(command, Tab, Space, -1)

	for _, s := range strings.Split(command, Space) {
		if s != "" {
			parsedCommand = append(parsedCommand, s)
		}
	}

	return parsedCommand
}

// processCommand process each command
func processCommand(p *Parking, command []string, first bool) (*Resp, error) {
	if len(command) == 0 {
		return nil, errors.New("Invalid input")
	}
	switch Input(command[0]) {
	case CreateParking:
		if !first {
			return &Resp{command: NotFisrt}, nil
		}
		maxSlots, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err.Error())
		}
		pl := NewParkingLot(maxSlots)
		p.parkingLot = pl
		return NewResponse(pl), nil
	case Park:
		return p.Park(command[1], command[2])
	case Status:
		return p.All()
	case Leave:
		slotPosition, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err.Error())
		}
		return p.LeaveByPosition(slotPosition)
	case CarRegNoWithdColour:
		return p.FindAllByColor(command[1], CarRegNoWithdColour)
	case SlotWithColour:
		return p.FindAllByColor(command[1], SlotWithColour)
	case SlotWithRegNo:
		return p.FindByRegistrationNumber(command[1])
	default:
		return nil, errors.New("Invalid input")
	}
}
