package parking

// Input is a typecast of string for input for commands
type Input string

const (
	CreateParking       Input = "create_parking_lot"
	Park                Input = "park"
	Leave               Input = "leave"
	Status              Input = "status"
	CarRegNoWithdColour Input = "registration_numbers_for_cars_with_colour"
	SlotWithColour      Input = "slot_numbers_for_cars_with_colour"
	SlotWithRegNo       Input = "slot_number_for_registration_number"
	NotFisrt            Input = "create_parking_lot_again"
)

const (
	Space                = " "
	Tab                  = "\t"
	MaxSlotReached       = "Sorry, parking lot is full"
	NoCarsParked         = "No cars parked"
	CarNotFound          = "Not found"
	CarWithColorNotFound = "Car with specified colour not found"
)
