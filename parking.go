package parking

type Parking struct {
	parkingLot *ParkingLot
}

func NewParking(maxSlots int) *Parking {
	pl := NewParkingLot(maxSlots)
	return &Parking{
		pl,
	}
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

func (p *Parking) FindAllByColor(color string, command Input) (*Resp, error) {
	slots, err := p.parkingLot.FindAllByColor(color)
	sResponse := &Resp{
		slots:   slots,
		command: command,
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
