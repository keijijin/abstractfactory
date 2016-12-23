package abstractfactory

// SportMotorbike is struct for SportMotorbike
type SportMotorbike struct{}

// GetWheels returns amount of wheels
func (s *SportMotorbike) GetWheels() int {
	return 2
}

// GetWSeats returns amount of seats
func (s *SportMotorbike) GetWSeats() int {
	return 1
}

// GetType returns type of motorbike
func (s *SportMotorbike) GetType() int {
	return SportMotorbikeType
}
