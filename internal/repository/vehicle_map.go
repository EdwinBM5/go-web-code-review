package repository

import (
	"app/internal"
	"fmt"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (r *VehicleMap) FindAll() (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	// copy db
	for key, value := range r.db {
		v[key] = value
	}

	return
}

// FindByColorAndYear is a method that returns a map of vehicles that match the color and year
func (r *VehicleMap) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	return
}

// FindByDimensions is a method that returns a map of vehicles that match the dimensions
func (r *VehicleMap) FindByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)
	fmt.Println(minLength, maxLength, minWidth, maxWidth)

	for key, value := range r.db {
		if value.Length >= minLength && value.Length <= maxLength && value.Width >= minWidth && value.Width <= maxWidth {
			v[key] = value
		}

		fmt.Println(value.Length, value.Width)
	}

	if len(v) == 0 {
		err = internal.ErrorVehicleNotFound
		return
	}

	return
}
