package abstractfactory

// LuxuryCar is struct for LuxuryCar
type LuxuryCar struct{}

// GetDoors returns amount of doors per a car
func (l *LuxuryCar) GetDoors() int {
	return 4
}

// GetWheels returns amount of Wheels
func (l *LuxuryCar) GetWheels() int {
	return 4
}

// GetSeats returns amount of seats
func (l *LuxuryCar) GetSeats() int {
	return 5
}
