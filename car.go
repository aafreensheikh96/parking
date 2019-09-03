package parking

// stores the struct of the car
type car struct {
	RegNo  string
	Colour string
}

// NewCar returns a new car onject
func NewCar(regNo, colour string) *car {
	return &car{
		RegNo:  regNo,
		Colour: colour,
	}
}
