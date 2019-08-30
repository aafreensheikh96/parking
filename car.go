package parking

type Car struct {
	RegNo  string
	Colour string
}

func NewCar(regNo, colour string) *Car {
	return &Car{
		RegNo:  regNo,
		Colour: colour,
	}
}
