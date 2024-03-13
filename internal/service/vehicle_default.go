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

// FindByBrandAndRangeYear is a method that returns a map of vehicles that match the brand and range year
func (s *VehicleDefault) FindByBrandAndRangeYear(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByBrandAndRangeYear(brand, startYear, endYear)
	return
}

// FindAverageSpeedByBrand is a method that returns a value of average speed by brand
func (s *VehicleDefault) FindAverageSpeedByBrand(brand string) (avgSpeed float64, err error) {
	avgSpeed, err = s.rp.FindAverageSpeedByBrand(brand)
	return
}

// FindByFuelType is a method that returns a map of vehicles that match the fuel type
func (s *VehicleDefault) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByFuelType(fuelType)
	return
}

// Delete is a method that deletes a vehicle
func (s *VehicleDefault) Delete(id int) (err error) {
	err = s.rp.Delete(id)
	return
}

// FindByTransmissionType is a method that returns a map of vehicles that match the transmission type
func (s *VehicleDefault) FindByTransmissionType(transmissionType string) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByTransmissionType(transmissionType)
	return
}

// FindAverageCapacityByBrand is a method that returns a value of average person capacity by brand
func (s *VehicleDefault) FindAverageCapacityByBrand(brand string) (avgCapacity float64, err error) {
	avgCapacity, err = s.rp.FindAverageCapacityByBrand(brand)
	return
}

// FindByDimensions is a method that returns a map of vehicles that match the dimensions
func (s *VehicleDefault) FindByDimensions(minHeight float64, maxHeight float64, minWidth float64, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v, err = s.rp.FindByDimensions(minHeight, maxHeight, minWidth, maxWidth)
	return
}
