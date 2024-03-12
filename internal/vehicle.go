package internal

import "errors"

// Dimensions is a struct that represents a dimension in 3d
type Dimensions struct {
	// Height is the height of the dimension
	Height float64
	// Length is the length of the dimension
	Length float64
	// Width is the width of the dimension
	Width float64
}

// VehicleAttributes is a struct that represents the attributes of a vehicle
type VehicleAttributes struct {
	// Brand is the brand of the vehicle
	Brand string
	// Model is the model of the vehicle
	Model string
	// Registration is the registration of the vehicle
	Registration string
	// Color is the color of the vehicle
	Color string
	// FabricationYear is the fabrication year of the vehicle
	FabricationYear int
	// Capacity is the capacity of people of the vehicle
	Capacity int
	// MaxSpeed is the maximum speed of the vehicle
	MaxSpeed float64
	// FuelType is the fuel type of the vehicle
	FuelType string
	// Transmission is the transmission of the vehicle
	Transmission string
	// Weight is the weight of the vehicle
	Weight float64
	// Dimensions is the dimensions of the vehicle
	Dimensions
}

// Vehicle is a struct that represents a vehicle
type Vehicle struct {
	// Id is the unique identifier of the vehicle
	Id int

	// VehicleAttribue is the attributes of a vehicle
	VehicleAttributes
}

// Errors in endpoints
var (
	// Error in most vehicle endpoints
	ErrorVehicleNotFound    = errors.New("Vehicle(s) not found")
	ErrorInternalServer     = errors.New("Internal server error")
	ErrorInvalidBodyRequest = errors.New("Invalid body request")
	// Error in input/vehicle attributes
	ErrorInvalidYear           = errors.New("Year must be a number")
	ErrorInvalidColorAndYear   = errors.New("Color and year are required")
	ErrorInvalidDimension      = errors.New("Invalid dimensions")
	ErrorInvalidHeightAndWidth = errors.New("Length and width are required")
	ErrorVehicleAlreadyExists  = errors.New("Vehicle already exists")
)
