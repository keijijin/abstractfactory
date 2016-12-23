package abstractfactory

// CruiseMotorbike is struct for CruiseMotorbike
type CruiseMotorbike struct{}

// GetWheels returns amount of wheels
func (c *CruiseMotorbike) GetWheels() int {
	return 2
}

// GetWSeats returns amount of seats
func (c *CruiseMotorbike) GetWSeats() int {
	return 1
}

// GetType returns type of motorbike
func (c *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}
