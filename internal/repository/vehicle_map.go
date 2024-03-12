package repository

import (
	"app/internal"
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

// Create is a method that creates a new vehicle
func (r *VehicleMap) Create(v *internal.Vehicle) (err error) {
	// check if vehicle already exists
	for _, value := range r.db {
		if value.Id == v.Id || value.Registration == v.Registration {
			err = internal.ErrorVehicleAlreadyExists
			return
		}
	}

	v.Id = len(r.db) + 1

	// add vehicle to db
	r.db[v.Id] = *v

	return
}

// FindByColorAndYear is a method that returns a map of vehicles that match the color and year
func (r *VehicleMap) FindByColorAndYear(color string, year int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Color == color && value.FabricationYear == year {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrorVehicleNotFound
	}

	return
}

// FindByBrandAndRangeYear is a method that returns a map of vehicles that match the brand and range year
func (r *VehicleMap) FindByBrandAndRangeYear(brand string, startYear int, endYear int) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Brand == brand && value.FabricationYear >= startYear && value.FabricationYear <= endYear {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrorVehicleNotFound
	}

	return
}

// FindAverageSpeedByBrand is a method that returns a value of average speed by brand
func (r *VehicleMap) FindAverageSpeedByBrand(brand string) (avgSpeed float64, err error) {
	var totalSpeed float64
	var brandCount int

	for _, value := range r.db {
		if value.Brand == brand {
			totalSpeed += value.MaxSpeed
			brandCount++
		}
	}

	if brandCount == 0 {
		err = internal.ErrorVehicleNotFound
		return
	}

	avgSpeed = totalSpeed / float64(brandCount)

	return
}

// FindByDimensions is a method that returns a map of vehicles that match the dimensions
func (r *VehicleMap) FindByDimensions(minHeight float64, maxHeight float64, minWidth float64, maxWidth float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Height >= minHeight && value.Height <= maxHeight && value.Width >= minWidth && value.Width <= maxWidth {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrorVehicleNotFound
		return
	}

	return
}
