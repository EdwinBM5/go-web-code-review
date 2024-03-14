package repository

import (
	"app/internal"
	"strings"
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

// FindByID is a method that returns a vehicle by ID
func (r *VehicleMap) FindByID(id int) (v internal.Vehicle, err error) {
	if _, ok := r.db[id]; !ok {
		err = internal.ErrorVehicleNotFound
		return
	}

	v = r.db[id]

	return
}

// Create is a method that creates a new vehicle
func (r *VehicleMap) Create(v *internal.Vehicle) (err error) {
	// check if vehicle already exists
	if v.Id == 0 {
		// generate new ID
		v.Id = len(r.db) + 1
	}

	for _, value := range r.db {
		if value.Id == v.Id || value.Registration == v.Registration {
			err = internal.ErrorVehicleAlreadyExists
			return
		}
	}

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

// UpdateMaxSpeed is a method that updates the max speed of a vehicle
func (r *VehicleMap) UpdateMaxSpeed(id int, maxSpeed float64) (err error) {
	if maxSpeed < 0 || maxSpeed > 500 {
		err = internal.ErrorInvalidMaxSpeedRange
		return
	}

	for key, value := range r.db {
		if value.Id == id {
			value.MaxSpeed = maxSpeed
			r.db[key] = value
			return
		}
	}

	err = internal.ErrorVehicleNotFound
	return
}

// FindByFuelType is a method that returns a map of vehicles that match the fuel type
func (r *VehicleMap) FindByFuelType(fuelType string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.FuelType == fuelType {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrorVehicleNotFound
	}

	return
}

// Delete is a method that deletes a vehicle
func (r *VehicleMap) Delete(id int) (err error) {
	for key, value := range r.db {
		if value.Id == id {
			delete(r.db, key)
			return
		}
	}

	err = internal.ErrorVehicleNotFound
	return
}

// FindByTransmissionType is a method that returns a map of vehicles that match the transmission type
func (r *VehicleMap) FindByTransmissionType(transmissionType string) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Transmission == transmissionType {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrorVehicleNotFound
	}

	return
}

// UpdateFuelType is a method that updates the fuel type of a vehicle
func (r *VehicleMap) UpdateFuelType(id int, fuelType string) (err error) {
	for key, value := range r.db {
		if value.Id == id {
			switch strings.ToLower(fuelType) {
			case "gasoline", "diesel", "biodiesel", "gas":
				value.FuelType = fuelType
				r.db[key] = value
			default:
				err = internal.ErrorInvalidFuelTypeUpdate
			}

			return
		}
	}

	err = internal.ErrorVehicleNotFound

	return
}

// FindAverageCapacityByBrand is a method that returns a value of average person capacity by brand
func (r *VehicleMap) FindAverageCapacityByBrand(brand string) (avgCapacity float64, err error) {
	var totalCapacity float64
	var brandCount int

	for _, value := range r.db {
		if value.Brand == brand {
			totalCapacity += float64(value.Capacity)
			brandCount++
		}
	}

	if brandCount == 0 {
		err = internal.ErrorVehicleNotFound
		return
	}

	avgCapacity = totalCapacity / float64(brandCount)

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
	}

	return
}

// FindByWeightRange is a method that returns a map of vehicles that match the weight range
func (r *VehicleMap) FindByWeightRange(minWeight float64, maxWeight float64) (v map[int]internal.Vehicle, err error) {
	v = make(map[int]internal.Vehicle)

	for key, value := range r.db {
		if value.Weight >= minWeight && value.Weight <= maxWeight {
			v[key] = value
		}
	}

	if len(v) == 0 {
		err = internal.ErrorVehicleNotFound
	}

	return
}
