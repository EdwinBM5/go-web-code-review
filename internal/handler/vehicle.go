package handler

import (
	"app/internal"
	"errors"
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

		// response
		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
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
		if err != nil {
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

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    v,
		})

	}
}

// Exercise twelve from code review
// GetByDimensions is a method that returns a handler for the route GET /vehicles/dimensions?length={length}&width={width}
func (h *VehicleDefault) GetByDimensions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		length := r.URL.Query().Get("length")
		width := r.URL.Query().Get("width")

		if length == "" || width == "" {
			response.Error(w, http.StatusBadRequest, internal.ErrorInvalidLengthAndWidth.Error())

			return
		}

		splitDimension := func(dimension string) (float64, float64, error) {
			split := strings.Split(dimension, "-")
			if len(split) != 2 {
				return 0, 0, internal.ErrorInvalidDimension
			}
			min, err := strconv.ParseFloat(split[0], 64)
			if err != nil {
				return 0, 0, err
			}
			max, err := strconv.ParseFloat(split[1], 64)
			if err != nil {
				return 0, 0, err
			}

			return min, max, nil
		}

		minLength, maxLength, err := splitDimension(length)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())

			return
		}

		minWidth, maxWidth, err := splitDimension(width)
		if err != nil {
			response.Error(w, http.StatusBadRequest, err.Error())

			return
		}

		v, err := h.sv.FindByDimensions(minLength, maxLength, minWidth, maxWidth)
		if err != nil {
			switch {
			case errors.Is(err, internal.ErrorVehicleNotFound):
				response.Error(w, http.StatusNotFound, err.Error())

			default:
				response.Error(w, http.StatusInternalServerError, internal.ErrorInternalServer.Error())
			}

			return
		}

		data := make(map[int]VehicleJSON)
		for key, value := range v {
			data[key] = VehicleJSON{
				ID:              value.Id,
				Brand:           value.Brand,
				Model:           value.Model,
				Registration:    value.Registration,
				Color:           value.Color,
				FabricationYear: value.FabricationYear,
				Capacity:        value.Capacity,
				MaxSpeed:        value.MaxSpeed,
				FuelType:        value.FuelType,
				Transmission:    value.Transmission,
				Weight:          value.Weight,
				Height:          value.Height,
				Length:          value.Length,
				Width:           value.Width,
			}
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})

	}
}
