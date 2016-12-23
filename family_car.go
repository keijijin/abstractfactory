package abstractfactory

// FamilyCar is struct for FamilyCar
type FamilyCar struct{}

// GetDoors returns amount of doors per a car
func (f *FamilyCar) GetDoors() int {
	return 5
}

// GetWheels returns amount of Wheels
func (f *FamilyCar) GetWheels() int {
	return 4
}

// GetSeats returns amount of seats
func (f *FamilyCar) GetSeats() int {
	return 5
}
