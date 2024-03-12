package internal

// VehicleRepository is an interface that represents a vehicle repository
type VehicleRepository interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	// Create is a method that creates a new vehicle
	Create(v *Vehicle) (err error)

	// FindByColorAndYear is a method that returns a map of vehicles that match the color and year
	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)

	// FindByBrandAndRangeYear is a method that returns a map of vehicles that match the brand and range year
	FindByBrandAndRangeYear(brand string, startYear int, endYear int) (v map[int]Vehicle, err error)

	// FindAverageSpeedByBrand is a method that returns a map of vehicles that match the average speed and brand
	FindAverageSpeedByBrand(brand string) (avgSpeed float64, err error)

	// FindByFuelType is a method that returns a map of vehicles that match the fuel type
	FindByFuelType(fuelType string) (v map[int]Vehicle, err error)

	// FindByTransmissionType is a method that returns a map of vehicles that match the transmission type
	FindByTransmissionType(transmissionType string) (v map[int]Vehicle, err error)

	// FindAverageCapacityByBrand is a method that returns a map of vehicles that match the average person capacity and brand
	FindAverageCapacityByBrand(brand string) (avgCapacity float64, err error)

	// FindByDimensions is a method that returns a map of vehicles that match the dimensions
	FindByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) (v map[int]Vehicle, err error)
}
