package abstractfactory

import (
	"errors"
	"fmt"
)

// VehicleFactory is interface for VehicleFactory
type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

// const for vehicle type
const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

// GetVehicleFactory returns object as vehicle type
func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))
	}
}

// const for car type
const (
	LuxuryCarType = 1
	FamilyCarType = 2
)

// CarFactory is struct for CarFactory
type CarFactory struct{}

// GetVehicle returns object for car type
func (c *CarFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}

// const for Motorbike Type
const (
	SportMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

// MotorbikeFactory is struct for MotorbikeFactory
type MotorbikeFactory struct{}

// GetVehicle returns object as type
func (m *MotorbikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
		//return new(SportMotorbike), nil
	case SportMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("Vehicle of type %d not recognized\n", v))
	}
}
