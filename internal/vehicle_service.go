package internal

// VehicleService is an interface that represents a vehicle service
type VehicleService interface {
	// FindAll is a method that returns a map of all vehicles
	FindAll() (v map[int]Vehicle, err error)

	// Create is a method that creates a new vehicle
	Create(v *Vehicle) (err error)

	// FindByColorAndYear is a method that returns a map of vehicles that match the color and year
	FindByColorAndYear(color string, year int) (v map[int]Vehicle, err error)

	// FindByDimensions is a method that returns a map of vehicles that match the dimensions
	FindByDimensions(minLength float64, maxLength float64, minWidth float64, maxWidth float64) (v map[int]Vehicle, err error)
}
