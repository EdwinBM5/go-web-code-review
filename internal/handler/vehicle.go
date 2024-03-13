package handler

import (
	"app/internal"
	"app/internal/tools"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// VehicleJSON is a struct that represents a vehicle in JSON format
type VehicleJSON struct {
	ID              int     `json:"id"`
	Brand           string  `json:"brand"`
	Model           string  `json:"model"`
	Registration    string  `json:"registration"`
	Color           string  `json:"color"`
	FabricationYear int     `json:"year"`
	Capacity        int     `json:"passengers"`
	MaxSpeed        float64 `json:"max_speed"`
	FuelType        string  `json:"fuel_type"`
	Transmission    string  `json:"transmission"`
	Weight          float64 `json:"weight"`
	Height          float64 `json:"height"`
	Length          float64 `json:"length"`
	Width           float64 `json:"width"`
}

// JSON is a method that returns a VehicleJSON from a Vehicle
func (v *VehicleJSON) JSON(vehicle internal.Vehicle) VehicleJSON {
	v.ID = vehicle.Id
	v.Brand = vehicle.Brand
	v.Model = vehicle.Model
	v.Registration = vehicle.Registration
	v.Color = vehicle.Color
	v.FabricationYear = vehicle.FabricationYear
	v.Capacity = vehicle.Capacity
	v.MaxSpeed = vehicle.MaxSpeed
	v.FuelType = vehicle.FuelType
	v.Transmission = vehicle.Transmission
	v.Weight = vehicle.Weight
	v.Height = vehicle.Height
	v.Length = vehicle.Length
	v.Width = vehicle.Width

	return *v
}

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(sv internal.VehicleService) *VehicleDefault {
	return &VehicleDefault{sv: sv}
}

// VehicleDefault is a struct with methods that represent handlers for vehicles
type VehicleDefault struct {
	// sv is the service that will be used by the handler
	sv internal.VehicleService
}

// GetAll is a method that returns a handler for the route GET /vehicles
func (h *VehicleDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		// - get all vehicles
		v, err := h.sv.FindAll()
		if err != nil {
			response.JSON(w, http.StatusInternalServerError, nil)
			return
		}

		// Prepare response in JSON format
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = (&VehicleJSON{}).JSON(value)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"count":   len(data),
			"message": "Success",
			"data":    data,
		})
	}
}

// Exercise one from code review
// Create is a method that returns a handler for the route POST /vehicles
func (h *VehicleDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// read body into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidBodyRequest.Error())

			return
		}

		// parse body to map to check required fields
		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidBodyRequest.Error())

			return
		}

		if err := tools.CheckFieldExistance(bodyMap, "brand", "model", "registration", "color", "year"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.Error(w, http.StatusBadRequest, fmt.Sprintf("%s is required", fieldError.Field))
				return
			}
		}

		// unmarshal bytes into VehicleJSON
		var v VehicleJSON
		if err := json.Unmarshal(bytes, &v); err != nil {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidBodyRequest.Error())

			return
		}

		// setup vehicle attributes
		vehicleDimension := internal.Dimensions{
			Length: v.Length,
			Width:  v.Width,
			Height: v.Height,
		}

		vehicleAttributes := internal.VehicleAttributes{
			Brand:           v.Brand,
			Model:           v.Model,
			Registration:    v.Registration,
			Color:           v.Color,
			FabricationYear: v.FabricationYear,
			Capacity:        v.Capacity,
			MaxSpeed:        v.MaxSpeed,
			FuelType:        v.FuelType,
			Transmission:    v.Transmission,
			Weight:          v.Weight,
			Dimensions:      vehicleDimension,
		}

		vehicle := internal.Vehicle{
			Id:                v.ID,
			VehicleAttributes: vehicleAttributes,
		}

		if err := h.sv.Create(&vehicle); err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleAlreadyExists):
				response.Error(w, http.StatusConflict, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		// Prepara JSON response
		data := (&VehicleJSON{}).JSON(vehicle)

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "Success",
			"data":    data,
		})
	}
}

// Exercise two from code review
// GetByColorAndYear is a method that returns a handler for the route GET /vehicles/color/{color}/year/{year}
func (h *VehicleDefault) GetByColorAndYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color := chi.URLParam(r, "color")
		year := chi.URLParam(r, "year")

		if color == "" || year == "" {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidColorAndYear.Error())

			return
		}

		fabricationYear, err := strconv.Atoi(year)
		if err != nil || fabricationYear < 0 {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidYear.Error())

			return
		}

		v, err := h.sv.FindByColorAndYear(color, fabricationYear)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleNotFound):
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		// Prepare response in JSON format
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = (&VehicleJSON{}).JSON(value)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"count":   len(data),
			"message": "Success",
			"data":    data,
		})

	}
}

// Exercise three from code review
// GetByBrandAndRangeYear is a method that returns a handler for the route GET /vehicles/brand/{brand}/between/{start_year}/{end_year}
func (h *VehicleDefault) GetByBrandAndRangeYear() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")
		startYear := chi.URLParam(r, "start_year")
		endYear := chi.URLParam(r, "end_year")

		if brand == "" || startYear == "" || endYear == "" {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidBrandAndRangeYear.Error())
			return
		}

		startYearInt, err := strconv.Atoi(startYear)
		if err != nil {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidYear.Error())
			return
		}

		endYearInt, err := strconv.Atoi(endYear)
		if err != nil {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidYear.Error())
			return
		}

		v, err := h.sv.FindByBrandAndRangeYear(brand, startYearInt, endYearInt)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleNotFound):
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		// Prepare response in JSON format
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = (&VehicleJSON{}).JSON(value)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"count":   len(data),
			"message": "Success",
			"data":    data,
		})

	}
}

// Exercise four from code review
// GetAverageSpeedByBrand is a method that returns a handler for the route GET /vehicles/average-speed/brand/{brand}
func (h *VehicleDefault) GetAverageSpeedByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		if brand == "" {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidBrand.Error())

			return
		}

		averageSpeed, err := h.sv.FindAverageSpeedByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleNotFound):
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Success",
			"data": map[string]any{
				"brand":         brand,
				"average_speed": averageSpeed,
			},
		})
	}
}

// Exercise seven from code review
// GetByFuelType is a method that returns a handler for the route GET /vehicles/fuel-type/{type}
func (h *VehicleDefault) GetByFuelType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fuelType := chi.URLParam(r, "type")
		if fuelType == "" {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidFuelType.Error())

			return
		}

		v, err := h.sv.FindByFuelType(fuelType)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleNotFound):
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		// Prepare response in JSON format
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = (&VehicleJSON{}).JSON(value)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"count":   len(data),
			"message": "Success",
			"data":    data,
		})
	}
}

// Exercise eight from code review
// Delete is a method that returns a handler for the route DELETE /vehicles/{id}
func (h *VehicleDefault) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidID.Error())

			return
		}

		idVehicle, err := strconv.Atoi(id)
		if err != nil {
			response.Error(w, http.StatusBadRequest, internal.ErrorParseID.Error())
			return
		}

		err = h.sv.Delete(idVehicle)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleNotFound):
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		response.JSON(w, http.StatusNoContent, nil)
	}
}

// Exercise nine from code review
// GetByTransmissionType is a method that returns a handler for the route GET /vehicles/transmission/{type}
func (h *VehicleDefault) GetByTransmissionType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transmissionType := chi.URLParam(r, "type")
		if transmissionType == "" {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidTransmissionType.Error())

			return
		}

		v, err := h.sv.FindByTransmissionType(transmissionType)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleNotFound):
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		// Prepare response in JSON format
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = (&VehicleJSON{}).JSON(value)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"count":   len(data),
			"message": "Success",
			"data":    data,
		})
	}
}

// Exercise eleven from code review
// GetAverageCapacityByBrand is a method that returns a handler for the route GET /vehicles/average-capacity/brand/{brand}
func (h *VehicleDefault) GetAverageCapacityByBrand() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brand := chi.URLParam(r, "brand")

		if brand == "" {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidBrand.Error())

			return
		}

		averageCapacity, err := h.sv.FindAverageCapacityByBrand(brand)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleNotFound):
				response.Error(w, http.StatusNotFound, err.Error())
			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Success",
			"data": map[string]any{
				"brand":            brand,
				"average_capacity": averageCapacity,
			},
		})
	}
}

// Exercise twelve from code review - No LENGTH uses width and height
// GetByDimensions is a method that returns a handler for the route GET /vehicles/dimensions?height={height}&width={width}
func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		height := r.URL.Query().Get("height")
		width := r.URL.Query().Get("width")

		if height == "" || width == "" {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidHeightAndWidth.Error())

			return
		}

		splitDimension := func(dimension string) (float64, float64, error) {
			split := strings.Split(dimension, "-")
			if len(split) != 2 {
				return 0, 0, internal.ErrorInvalidDimension
			}
			min, err := strconv.ParseFloat(split[0], 64)
			if err != nil {
				return 0, 0, internal.ErrorInvalidQueryParamFormat
			}
			max, err := strconv.ParseFloat(split[1], 64)
			if err != nil {
				return 0, 0, internal.ErrorInvalidQueryParamFormat
			}

			return min, max, nil
		}

		minHeight, maxHeight, err := splitDimension(height)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())

			return
		}

		minWidth, maxWidth, err := splitDimension(width)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())

			return
		}

		v, err := h.sv.FindByDimensions(minHeight, maxHeight, minWidth, maxWidth)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleNotFound):
				response.Error(w, http.StatusNotFound, err.Error())

			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		// Prepare response in JSON format
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = (&VehicleJSON{}).JSON(value)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"count":   len(data),
			"message": "Success",
			"data":    data,
		})
	}
}

// Exercise thirteen from code review
// GetByWeightRange is a method that returns a handler for the route GET /vehicles/weight?min={weight_min}&max={weight_max}
func (h *VehicleDefault) GetByWeightRange() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		minWeight := r.URL.Query().Get("min")
		maxWeight := r.URL.Query().Get("max")

		if minWeight == "" || maxWeight == "" {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidWeightRange.Error())

			return
		}

		minWeightFloat, err := strconv.ParseFloat(minWeight, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidQueryParamFormat.Error())

			return
		}

		maxWeightFloat, err := strconv.ParseFloat(maxWeight, 64)
		if err != nil {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidQueryParamFormat.Error())

			return
		}

		v, err := h.sv.FindByWeightRange(minWeightFloat, maxWeightFloat)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleNotFound):
				response.Error(w, http.StatusNotFound, err.Error())

			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		// Prepare response in JSON format
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = (&VehicleJSON{}).JSON(value)
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"count":   len(data),
			"message": "Success",
			"data":    data,
		})

	}
}
