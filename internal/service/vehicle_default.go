package service

import "app/internal"

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{rp: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	rp internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (s *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindAll()
	return
}

// Create is a method that creates a new vehicle
func (s *VehicleDefault) Create(v *internal.Vehicle) (err error) {
	err = s.rp.Create(v)
	return
}

// FindByColorAndYear is a method that returns a map of vehicles that match the color and year
func (s *VehicleDefault) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByColorAndYear(color, year)
	return
}

// FindByDimensions is a method that returns a map of vehicles that match the dimensions
func (s *VehicleDefault) FindByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByDimensions(minLength, maxLength, minWidth, maxWidth)
	return
}
