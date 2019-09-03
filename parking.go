package parking

// struct that hold the parking lot.
type parking struct {
	parkingLot *parkingLot
}

// NewParking creates new parking area which stores the parking lot
func NewParking(maxSlots int) *parking {
	pl := NewParkingLot(maxSlots)
	return &parking{
		pl,
	}
}

// Park is a wrapper over the `parkingLot.Park`
func (p *parking) Park(numberPlate, color string) (*resp, error) {
	s, err := p.parkingLot.Park(numberPlate, color)
	r := &resp{
		slots: []*slot{
			s,
		},
		command: Park,
	}

	return r, err
}

// LeaveByPosition is a wrapper over the `parkingLot.LeaveByPosition`
func (p *parking) LeaveByPosition(position int) (*resp, error) {
	s, err := p.parkingLot.LeaveByPosition(position)
	r := &resp{
		slots: []*slot{
			s,
		},
		command: Leave,
	}

	return r, err
}

// FindByRegistrationNumber is a wrapper over the `parkingLot.FindByRegistrationNumber`
func (p *parking) FindByRegistrationNumber(numberPlate string) (*resp, error) {
	s, err := p.parkingLot.FindByRegistrationNumber(numberPlate)
	r := &resp{
		slots:   []*slot{s},
		command: SlotWithRegNo,
	}

	return r, err
}

// FindAllByColor is a wrapper over the `parkingLot.FindAllByColor`
func (p *parking) FindAllByColor(color string, command Input) (*resp, error) {
	slots, err := p.parkingLot.FindAllByColor(color)
	r := &resp{
		slots:   slots,
		command: command,
	}
	return r, err
}

// All is a wrapper over the `parkingLot.AllSlots`
func (p *parking) All() (*resp, error) {
	slots, err := p.parkingLot.AllSlots()
	return &resp{
		slots:   slots,
		command: Status,
	}, err
}
